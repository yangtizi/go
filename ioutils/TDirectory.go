package ioutils

import (
	"io"
	"os"
	"path/filepath"
)

// Directory .
var Directory = &TDirectory{}

// TDirectory 文件夹相关的函数在这里
type TDirectory struct{}

//

// TStringDynArray .
type TStringDynArray []string

// TSearchOption .
type TSearchOption int

// type TFileAttribute = (faReadOnly, faHidden, faSystem, faDirectory, faArchive,
// faDevice, faNormal, faTemporary, faSparseFile, faReparsePoint, faCompressed,
// faOffline, faNotContentIndexed, faEncrypted, faSymLink)

const (
	_ TSearchOption = -1 + iota
	//TopDirectoryOnly  当前目录
	TopDirectoryOnly
	// AllDirectories 所有目录
	AllDirectories
)

// https://www.cnblogs.com/del/archive/2009/10/16/1584660.html
// Walk 的具体实现在这里
// func (*TDirectory)getAllFile(pathname string, s TStringDynArray) (TStringDynArray, error) {
// 	rd, err := ioutil.ReadDir(pathname)
// 	if err != nil {
// 		fmt.Println("read dir fail:", err)
// 		return s, err
// 	}
// 	for _, fi := range rd {
// 		if fi.IsDir() {
// 			fullDir := pathname + fi.Name() + "/"
// 			s, err = getAllFile(fullDir, s)
// 			if err != nil {
// 				fmt.Println("read dir fail:", err)
// 				return s, err
// 			}
// 		} else {
// 			fullName := pathname + fi.Name()
// 			s = append(s, fullName)
// 		}
// 	}
// 	return s, nil
// }

// GetFilesAndDir 获取文件夹和dir
func (*TDirectory) GetFilesAndDir(strPath string) (TStringDynArray, TStringDynArray, error) {
	var files TStringDynArray
	var dirs TStringDynArray
	filepath.Walk(strPath, func(strFilename string, info os.FileInfo, err error) error {
		strFilename = filepath.ToSlash(strFilename)
		if err != nil {
			return err
		}
		if info.IsDir() {
			dirs = append(dirs, strFilename)
			return err
		}

		files = append(files, strFilename)
		return err
	})

	// files, _ := getAllFile(strPath, files)
	return files, dirs, nil
}

// GetFiles 使用通配符(暂时不支持通配符功能)
func (*TDirectory) GetFiles(strPath, strSearchPattern string, SearchOption TSearchOption) (TStringDynArray, error) {
	var files TStringDynArray
	filepath.Walk(strPath, func(strFilename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return err
		}

		strFilename = filepath.ToSlash(strFilename)

		// 影响效率的通配符
		// b, err := filepath.Match(strSearchPattern, filepath.Base(strFilename))
		// if err != nil {
		// 	return err
		// }

		// if !b {
		//     return err
		// }

		files = append(files, strFilename)

		return err
	})

	// files, _ := getAllFile(strPath, files)
	return files, nil
}

// CreateDirectory 建立新目录
func (m *TDirectory) CreateDirectory(path string) error {
	if !m.Exists(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

// Exists 判断文件夹是否存在
func (*TDirectory) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsEmpty 判断文件夹是否为空
func (*TDirectory) IsEmpty(strPath string) bool {
	fd, err := os.Open("d:/aaa/")
	if err != nil {
		return false
	}
	names, _ := fd.Readdirnames(1)
	return len(names) == 0
}

/*
	src := "./src"
	dst := "./dst"
	err := CopyDirFiles(src, dst)
	if err != nil {
		fmt.Println(err)
	}
*/
// Copy 复制文件夹
func (*TDirectory) Copy(src, dst string) error {
	// This code uses the filepath.Walk function to recursively traverse the directory specified by src. For each file in the directory, the CopyDirFiles function opens the source file and creates a destination file. The contents of the source file are then copied to the destination file using the io.Copy function. If an error occurs during the copy, the error is returned to the caller.
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		srcPath := path
		dstPath := filepath.Join(dst, filepath.Base(path))

		if info.IsDir() {
			return nil
		}

		srcFile, err := os.Open(srcPath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			return err
		}

		return nil
	})
}

// Move 移动文件夹
func (*TDirectory) Move(src, dst string) error {
	return os.Rename(src, dst)
}

// Delete 删除文件夹, 第二个参数为 True 可删除非空文件夹
func (m *TDirectory) Delete(strPath string, bForceDelete bool) error {
	if bForceDelete {
		return os.RemoveAll(strPath)
	}

	return os.Remove(strPath)
}

// GetDirectoryRoot 获取目录的根盘符, 如: C:\
func (*TDirectory) GetDirectoryRoot(strPath string) string {
	return filepath.VolumeName(strPath)
}

// GetCurrentDirectory 获取当前目录
func (*TDirectory) GetCurrentDirectory() (string, error) {
	return os.Getwd()
}

// SetCurrentDirectory 设置当前目录
func (*TDirectory) SetCurrentDirectory(strPath string) error {
	return os.Chdir(strPath)
}

// GetLogicalDrives 获取驱动器列表; 下有举例
func (*TDirectory) GetLogicalDrives() []string {
	var drives []string
	return drives
}

/*
	fmt.Println("Name:", fi.Name())
	fmt.Println("Size:", fi.Size())
	fmt.Println("Mode:", fi.Mode())
	fmt.Println("ModTime:", fi.ModTime())
	fmt.Println("IsDir:", fi.IsDir())
*/
// GetAttributes 获取文件夹属性, 譬如只读、存档等; 下有举例
func (*TDirectory) GetAttributes(strPath string) (os.FileInfo, error) {
	fi, err := os.Stat(strPath)
	if err != nil {
		return nil, err
	}
	return fi, nil
}

// SetAttributes 设置文件夹属性; 下有举例
// mode := os.FileMode(0777)
// err := SetAttributes("./test", mode)
func (*TDirectory) SetAttributes(strPath string, mode os.FileMode) error {
	err := os.Chmod(strPath, mode)
	if err != nil {
		return err
	}
	return nil
}
