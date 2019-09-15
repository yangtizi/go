package tdirectory

import (
	"fmt"
	"io/ioutil"
	"os"
)

// import "github.com/yangtizi/go/ioutils"

// TStringDynArray .
type TStringDynArray []string

// TSearchOption .
type TSearchOption int

const (
	_ TSearchOption = -1 + iota
	//TopDirectoryOnly  当前目录
	TopDirectoryOnly
	// AllDirectories 所有目录
	AllDirectories
)

// https://www.cnblogs.com/del/archive/2009/10/16/1584660.html

func getAllFile(pathname string, s TStringDynArray) (TStringDynArray, error) {
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := pathname + fi.Name() + "/"
			s, err = getAllFile(fullDir, s)
			if err != nil {
				fmt.Println("read dir fail:", err)
				return s, err
			}
		} else {
			fullName := pathname + fi.Name()
			s = append(s, fullName)
		}
	}
	return s, nil
}

// GetFiles 使用通配符(暂时不支持通配符功能)
func GetFiles(strPath, strSearchPattern string, SearchOption TSearchOption) (TStringDynArray, error) {
	var files TStringDynArray
	files, err := getAllFile(strPath, files)

	return files, err
}

// CreateDirectory 建立新目录
func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// Exists 判断文件夹是否存在
func Exists(path string) bool {
	return false
}

// IsEmpty 判断文件夹是否为空
func IsEmpty() {

}

// Copy 复制文件夹
func Copy(dstName, srcName string) {
}

// Move 移动文件夹
func Move() {

}

// Delete 删除文件夹, 第二个参数为 True 可删除非空文件夹
func Delete() {

}

// GetDirectoryRoot 获取目录的根盘符, 如: C:\
func GetDirectoryRoot() {

}

// GetCurrentDirectory 获取当前目录
func GetCurrentDirectory() {

}

// SetCurrentDirectory 设置当前目录
func SetCurrentDirectory() {

}

// GetLogicalDrives 获取驱动器列表; 下有举例
func GetLogicalDrives() {

}

// GetAttributes 获取文件夹属性, 譬如只读、存档等; 下有举例
func GetAttributes() {

}

// SetAttributes 设置文件夹属性; 下有举例
func SetAttributes() {

}
