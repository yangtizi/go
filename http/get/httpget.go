package get

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/yangtizi/go/log/zaplog"
)

// Bytes 函数
func Bytes(strURL string) ([]byte, error) {
	resp, err := http.Get(strURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zaplog.Println(err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return body, errors.New(resp.Status)
	}

	return body, nil
}
