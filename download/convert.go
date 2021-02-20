package download

import (
	"fmt"
	"time"
)

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

var nDownloaded int64   // 已下载内容
var bOpenSurfCount bool // 流量统计
var strSpeed string     // 实时网速

// OpenSurfCount 打开流量统计
func OpenSurfCount() {
	if bOpenSurfCount {
		return
	}
	bOpenSurfCount = true

	nDownloaded = 0

	go func() {
		for {
			time.Sleep(1 * time.Second)
			strSpeed = SizeToBytes(nDownloaded)
			nDownloaded = 0
			if !bOpenSurfCount {
				return
			}
		}
	}()

}

// CloseSurfCount 关闭流量统计
func CloseSurfCount() {
	bOpenSurfCount = false
}

// GetSpeed 获取速度
func GetSpeed() string {
	return strSpeed
}
