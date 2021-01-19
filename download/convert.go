package download

import "fmt"

// SizeToBytes 文件大小转换成可阅读的类型 nSize 尺寸
func SizeToBytes(nSize int64) string {
	if nSize > 1048576 { // 1024 * 1024
		return fmt.Sprintf("%.2f MB", float64(nSize)/1048576.0)
	}

	if nSize > 1024 {
		return fmt.Sprintf("%.2f KB", float64(nSize)/1024.0)
	}

	return fmt.Sprintf("%d B", nSize)
}
