package post

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/yangtizi/go/log/zaplog"
)

// JSON (strURL 网址,  strJson POST内容)
func JSON(strURL, strJSON string) string {
	resp, err := http.Post(
		strURL,
		"application/json",
		strings.NewReader(strJSON))

	if err != nil {
		log.Println(err)
		return string(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return string(err.Error())
	}

	// log.Println(string(body))
	return string(body)
}

// Bytes 这里是发送BUFF内容
func Bytes(strURL string, buf []byte) ([]byte, error) {
	resp, err := http.Post(
		strURL,
		"",
		bytes.NewReader(buf))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	// log.Println(string(body))
	return body, err
}
