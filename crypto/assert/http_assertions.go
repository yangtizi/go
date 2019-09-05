package assert

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// httpCode 是返回响应的HTTP代码的帮助程序。它返回-1和
//如果生成新请求失败，则会出现错误。
func httpCode(handler http.HandlerFunc, method, url string, values url.Values) (int, error) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return -1, err
	}
	req.URL.RawQuery = values.Encode()
	handler(w, req)
	return w.Code, nil
}

// HTTPSuccess 断言指定的处理程序返回成功状态代码。
//
//  assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
//返回断言是否成功（true）或不成功（false）。
func HTTPSuccess(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	code, err := httpCode(handler, method, url, values)
	if err != nil {
		Fail(t, fmt.Sprintf("Failed to build test request, got error: %s", err))
		return false
	}

	isSuccessCode := code >= http.StatusOK && code <= http.StatusPartialContent
	if !isSuccessCode {
		Fail(t, fmt.Sprintf("Expected HTTP success status code for %q but received %d", url+"?"+values.Encode(), code))
	}

	return isSuccessCode
}

// HTTPRedirect 断言指定的处理程序返回重定向状态代码。
//
//  assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
//返回断言是否成功（true）或不成功（false）。
func HTTPRedirect(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	code, err := httpCode(handler, method, url, values)
	if err != nil {
		Fail(t, fmt.Sprintf("Failed to build test request, got error: %s", err))
		return false
	}

	isRedirectCode := code >= http.StatusMultipleChoices && code <= http.StatusTemporaryRedirect
	if !isRedirectCode {
		Fail(t, fmt.Sprintf("Expected HTTP redirect status code for %q but received %d", url+"?"+values.Encode(), code))
	}

	return isRedirectCode
}

// HTTPError 断言指定的处理程序返回错误状态代码。
//
//  assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
//返回断言是否成功（true）或不成功（false）。
func HTTPError(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	code, err := httpCode(handler, method, url, values)
	if err != nil {
		Fail(t, fmt.Sprintf("Failed to build test request, got error: %s", err))
		return false
	}

	isErrorCode := code >= http.StatusBadRequest
	if !isErrorCode {
		Fail(t, fmt.Sprintf("Expected HTTP error status code for %q but received %d", url+"?"+values.Encode(), code))
	}

	return isErrorCode
}

// HTTPBody 是返回响应的HTTP主体的帮助程序。它回来了
//如果生成新请求失败，则为空字符串。
func HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) string {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url+"?"+values.Encode(), nil)
	if err != nil {
		return ""
	}
	handler(w, req)
	return w.Body.String()
}

// HTTPBodyContains 断言指定的处理程序返回
//包含字符串的正文。
//
//  assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "我很幸运")
//
//返回断言是否成功（true）或不成功（false）。
func HTTPBodyContains(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	body := HTTPBody(handler, method, url, values)

	contains := strings.Contains(body, fmt.Sprint(str))
	if !contains {
		Fail(t, fmt.Sprintf("Expected response body for \"%s\" to contain \"%s\" but found \"%s\"", url+"?"+values.Encode(), str, body))
	}

	return contains
}

// HTTPBodyNotContains 断言指定的处理程序返回
//不包含字符串的正文。
//
//assert.httpbodynotcontains（t，myhandler，“get”，“www.google.com”，nil，“我很幸运”）。
//
//返回断言是否成功（true）或不成功（false）。
func HTTPBodyNotContains(t TestingT, handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	body := HTTPBody(handler, method, url, values)

	contains := strings.Contains(body, fmt.Sprint(str))
	if contains {
		Fail(t, fmt.Sprintf("Expected response body for \"%s\" to NOT contain \"%s\" but found \"%s\"", url+"?"+values.Encode(), str, body))
	}

	return !contains
}
