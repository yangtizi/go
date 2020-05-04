package post

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	log "github.com/yangtizi/go/log/zaplog"

	"golang.org/x/net/proxy"
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

	// log.Println(string(body))PEv0nLUQnWjFcCiv PEv0nLUQnWjFcCiv PEv0nLUQnWjFcCiv PEv0nLUQnWjFcCiv
	return body, err
}

// HTTPProxy 使用HTTP代理
func HTTPProxy(strURL string, buf []byte, strProxy string, strUser string, strPassword string) ([]byte, error) {

	// 配置代理
	urli := url.URL{}
	urlproxy, _ := urli.Parse(strProxy)
	// 拿出代理
	client := &http.Client{}

	if len(strProxy) > 7 {
		log.Println("使用代理", urlproxy)
		client.Transport = &http.Transport{Proxy: http.ProxyURL(urlproxy)}
	} else {
		log.Println("不是用代理", strProxy)
	}

	// 请求
	req, err := http.NewRequest("POST", strURL, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	// 设置头
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "close")
	req.Header.Set("Content-type", "application/octet-stream")
	req.Header.Set("User-Agent", "MicroMessenger Client")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取 respon
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

// Sock5Proxy Sock5代理
func Sock5Proxy(strURL string, buf []byte, strProxy string, strUser string, strPassword string) ([]byte, error) {
	auth := &proxy.Auth{
		User:     strUser,
		Password: strPassword,
	}

	if strURL == "" && strPassword == "" {
		auth = nil
	}

	dialer, err := proxy.SOCKS5("tcp", strProxy, auth, proxy.Direct)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	if len(strProxy) > 7 {
		log.Println("使用代理", strProxy)
		client.Transport = &http.Transport{Dial: dialer.Dial}
	} else {
		log.Println("不是用代理", strProxy)
	}

	// 请求
	req, err := http.NewRequest("POST", strURL, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	// 设置头
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "close")
	req.Header.Set("Content-type", "application/octet-stream")
	req.Header.Set("User-Agent", "MicroMessenger Client")

	// 执行请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取 respon
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
