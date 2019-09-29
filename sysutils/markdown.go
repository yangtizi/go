package sysutils

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// MakedownToHTML 转换成普通的
func MakedownToHTML(md []byte) []byte {	
	u := blackfriday.Run(md)
	html := bluemonday.UGCPolicy().SanitizeBytes(u)

	return html
}
