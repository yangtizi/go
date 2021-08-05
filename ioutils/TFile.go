package ioutils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// File .
var File = &TFile{}

// TFile 文件相关的函数在里面
type TFile struct{}

// Exists 判断指定的文件是否存在
func (*TFile) Exists(f string) bool {
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
func (*TFile) Copy(src, dst string) (int64, error) {
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
func (*TFile) Move(src, dst string) error {
	return os.Rename(src, dst)
}

// Delete 删除文件
func (*TFile) Delete(s string) error {
	return os.Remove(s)
}

// OpenRead 读取一个文件
func (*TFile) OpenRead(s string) ([]byte, error) {
	b, err := ioutil.ReadFile(s)
	return b, err
}

// Create 创建一个新文件， 如果文件存在，那么会直接被清空掉内容
func (*TFile) Create(s string, b []byte) error {
	return ioutil.WriteFile(s, b, 0666)

	//  ioutil 包里的代码
	// 	func WriteFile(filename string, data []byte, perm os.FileMode) error {
	// 	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	_, err = f.Write(data)
	// 	if err1 := f.Close(); err == nil {
	// 		err = err1
	// 	}
	// 	return err
	// }

	// os 包里的实现代码
	// f, err := os.Create(s) // OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	// if err != nil {
	// 	return err
	// }
	// f.Write(b)
	// err = f.Close()
	// return err
}

// JSON 获取文件的json值
func (m *TFile) JSON(s string, v interface{}) error {
	data, err := m.OpenRead(s)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// MD5 计算MD5值
func (m *TFile) MD5(s string) (string, error) {
	data, err := m.OpenRead(s)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", md5.Sum(data)), nil
}
