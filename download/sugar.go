package download

import "os"

// 下载

func ResumeBreakPoint(strFile string, strURL string, nTotal int64, cb *TCallback) error {
	p := NewResumeBreakPoint()

	p.SetCallback(cb)
	p.SetTotal(nTotal)
	err := p.OpenFile(strFile, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer p.Close()
	return p.Download(strURL)
}
