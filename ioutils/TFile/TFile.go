package tfile

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Exists 判断指定的文件是否存在
func Exists(f string) bool {
	_, err := os.Stat(f) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// Copy 复制文件
func Copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}

	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// Move 移动文件
func Move(src, dst string) error {
	return os.Rename(src, dst)
}

// Delete 删除文件
func Delete(s string) error {
	return os.Remove(s)
}

// OpenRead 读取一个文件
func OpenRead(s string) ([]byte, error) {
	b, err := ioutil.ReadFile(s)
	return b, err
}
