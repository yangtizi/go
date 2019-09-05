package zlib

// Zlib 压缩与解析库
import (
	"bytes"
	z "compress/zlib"
	"io"
)

// Compress 进行zlib压缩
func Compress(src []byte) []byte {
	var in bytes.Buffer
	w := z.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

// UnCompress 进行zlib解压缩
func UnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := z.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}
