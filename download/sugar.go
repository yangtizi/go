package download

import (
	"os"

	"github.com/yangtizi/go/ioutils"
)

// ResumeBreakPoint 通过断点续传下载文件
// strFile 保存的文件名
// strURL 下载地址
// nTotal 文件大小
func ResumeBreakPoint(strFile string, strURL string, nTotal int64, cb TCallback) error {
	p := NewResumeBreakPoint()
	p.SetCallback(cb)
	p.SetTotal(nTotal)

	// 下载的目录名
	strPath := ioutils.Path.GetDirectoryName(strFile)
	// 先创建目录
	ioutils.Directory.CreateDirectory(strPath)

	// 创建文件
	err := p.OpenFile(strFile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer p.Close()
	return p.Download(strURL)
}
