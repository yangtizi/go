package tpath

import (
	"os"
	"path/filepath"
)

// GetTempPath 获取临时文件夹路径
func GetTempPath() {
}

// GetTempFileName 获取一个临时文件名
func GetTempFileName() {
}

// GetPathRoot 提取盘符, 如: c:\
func GetPathRoot() {
}

// GetDirectoryName 提取路径
func GetDirectoryName(path string) string {
	return filepath.Dir(path)
}

// GetFileName 提取文件名
func GetFileName() {
}

// GetExtension 提取扩展名
func GetExtension(path string) string {
	path = filepath.Base(path)
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[0:i]
		}
	}
	return ""
}

// GetFileNameWithoutExtension 提取无扩展名的文件名
func GetFileNameWithoutExtension() {
}

// ChangeExtension 更换扩展名
func ChangeExtension() {
}

// DriveExists 检查路径中的驱动器是否存在
func DriveExists() {
}

// GetFullPath 根据相对路径给出全路径
func GetFullPath() {
}

// HasExtension 判断是否有扩展名
func HasExtension() {
}

// IsPathRooted 判断是否是绝对路径
func IsPathRooted() {
}

// Combine 结合路径
func Combine() {
}

// GetRandomFileName 产生一个随机文件名
func GetRandomFileName() {
}

// GetGUIDFileName 用于产生一个唯一的文件名, 布尔参数决定名称中是否包含 -
func GetGUIDFileName() {
}

// IsValidPathChar 判断给定的字符是否能用于路径名
func IsValidPathChar() {
}

// IsValidFileNameChar 判断给定的字符是否能用于文件名
func IsValidFileNameChar() {
}
