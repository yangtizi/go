package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// 不要扩展名的文件
func noExt(path string) string {
	path = filepath.Base(path)
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[0:i]
		}
	}
	return ""
}

func init() {
	//设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// SetLogFilename 设置文件名称
func SetLogFilename(strFilename string) {
	outfile, err := os.OpenFile(strFilename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//设置log的输出文件，不设置log输出默认为stdout
	log.SetOutput(outfile)
}

// AutoLogFilename 自动设置Filename
func AutoLogFilename() {
	SetLogFilename(autologfilename())
}

func autologfilename() string {
	strPath := "./log/" + noExt(os.Args[0]) + "/"
	fmt.Println("保存日志", strPath)
	os.MkdirAll(strPath, os.ModePerm)

	strFilename := strPath + time.Now().Format("20060102") + ".log"
	return strFilename
}
