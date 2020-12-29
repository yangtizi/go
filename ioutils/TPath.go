package ioutils

import (
	"os"
	"path/filepath"

	"github.com/yangtizi/go/sysutils"
)

// Path .
var Path = &TPath{}

// TPath 路径相关的函数
type TPath struct{}

// GetTempPath 获取临时文件夹路径
func (*TPath) GetTempPath() string {
	return os.TempDir()
}

// GetTempFileName 获取一个临时文件名
func (*TPath) GetTempFileName() string {
	return sysutils.GenUUID()
}

// GetPathRoot 提取盘符, 如: c:\
func (*TPath) GetPathRoot(strPath string) string {
	return filepath.VolumeName(strPath)
}

// GetDirectoryName 提取路径
func (*TPath) GetDirectoryName(strPath string) string {
	return filepath.Dir(strPath)
}

// GetFileName 提取文件名
func (*TPath) GetFileName(strPath string) string {
	return filepath.Base(strPath)
}

// GetExtension 提取扩展名
func (*TPath) GetExtension(strPath string) string {
	return filepath.Ext(strPath)
}

// GetFileNameWithoutExtension 提取无扩展名的文件名
func (*TPath) GetFileNameWithoutExtension(strPath string) {

}

// ChangeExtension 更换扩展名
func (*TPath) ChangeExtension() {
}

// DriveExists 检查路径中的驱动器是否存在
func (*TPath) DriveExists() {
}

// GetFullPath 根据相对路径给出全路径
func (*TPath) GetFullPath(strPath string) string {
	strPath, _ = filepath.Abs(strPath)
	return filepath.ToSlash(strPath) + "/"
}

// HasExtension 判断是否有扩展名
func (*TPath) HasExtension() {
}

// IsPathRooted 判断是否是绝对路径
func (*TPath) IsPathRooted(strPath string) bool {
	return filepath.IsAbs(strPath)
}

// Combine 结合路径
func (*TPath) Combine(strPath string) string {
	return filepath.Clean(strPath)
}

// GetRandomFileName 产生一个随机文件名
func (*TPath) GetRandomFileName() string {
	return sysutils.GenUUID()
}

// GetGUIDFileName 用于产生一个唯一的文件名, 布尔参数决定名称中是否包含 -
func (*TPath) GetGUIDFileName() string {
	return sysutils.GenUUID()
}

// IsValidPathChar 判断给定的字符是否能用于路径名
func (*TPath) IsValidPathChar() {
}

// IsValidFileNameChar 判断给定的字符是否能用于文件名
func (*TPath) IsValidFileNameChar() {
}
