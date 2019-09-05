package jsonconfig

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"strings"

	log "github.com/yangtizi/go/log/zaplog"
)

//JSONParsing 解析json文件, 文件名, 解析的JSON结构体
func JSONParsing(strPath string, v interface{}) {
	//创建一个新的buff
	buf := new(bytes.Buffer)
	//打开文件
	f, err := os.Open(strPath)
	// f,err :=  io.ReadFile(strPath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	//处理注释
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadSlice('\n')
		if err != nil {
			if len(line) > 0 {
				buf.Write(line)
			}
			break
		}
		if !strings.HasPrefix(strings.TrimLeft(string(line), "\t"), "//") {
			buf.Write(line)
		}
	}
	//解析json

	json.Unmarshal([]byte(buf.Bytes()), v)
}
