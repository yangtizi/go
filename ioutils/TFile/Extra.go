package tfile

import (
	"encoding/json"

	"github.com/yangtizi/go/crypto/md5"
)

// JSON 获取文件的json值
func JSON(s string, v interface{}) error {
	data, err := OpenRead(s)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// MD5 计算MD5值
func MD5(s string) (string, error) {
	data, err := OpenRead(s)
	if err != nil {
		return "", err
	}

	return md5.BinToHex(data), nil
}
