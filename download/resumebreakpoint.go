package download

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// TCallback 回调函数
type TCallback func(nCurrent, nTotal int64) error

// TResumeBreakPoint 断点续传
type TResumeBreakPoint struct {
	f        *os.File  // 文件指针
	nCurrent int64     // 当前长度
	nTotal   int64     // 总长度
	cb       TCallback // 回调
}

// NewResumeBreakPoint 新建断点下载类
func NewResumeBreakPoint() *TResumeBreakPoint {
	p := &TResumeBreakPoint{}
	return p
}

// OpenFile 打开文件, 参考os.OpenFile
func (m *TResumeBreakPoint) OpenFile(name string, flag int, perm os.FileMode) error {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return err
	}

	fstat, err := f.Stat()
	if err != nil {
		defer f.Close()
		return err
	}

	m.f = f
	m.nCurrent = fstat.Size()
	f.Seek(m.nCurrent, 0)
	return nil
}

// Close 关闭
func (m *TResumeBreakPoint) Close() {
	if m.f != nil {
		defer m.f.Close()
	}
}

// Download 下载文件
func (m *TResumeBreakPoint) Download(strURL string) error {
	req, err := http.NewRequest(http.MethodGet, strURL, nil)
	if err != nil {
		return err
	}

	strRange := fmt.Sprintf("bytes=%d-", m.nCurrent)
	req.Header.Set("Range", strRange)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 206 && resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	_, err = io.Copy(m.f, io.TeeReader(resp.Body, m))
	return err
}

// Write 写入操作
func (m *TResumeBreakPoint) Write(p []byte) (int, error) {
	n := len(p)
	m.nCurrent += int64(n)
	if bOpenSurfCount {
		nDownloaded += int64(n)
	}
	if m.cb != nil {
		return n, m.cb(m.nCurrent, m.nTotal)
	}
	fmt.Printf("\r                                  ")
	fmt.Printf("\r %s / %s", SizeToBytes(m.nCurrent), SizeToBytes(m.nTotal))
	return n, nil
}

// SetCallback 设置Callback
func (m *TResumeBreakPoint) SetCallback(cb TCallback) {
	m.cb = cb
}

// SetTotal 设置总数
func (m *TResumeBreakPoint) SetTotal(nTotal int64) {
	m.nTotal = nTotal
}
