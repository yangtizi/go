package assert

import (
	"errors"
)

// AnError 是一个对测试有用的错误实例。如果代码不关心
//* 关于错误的详细信息，并且只需要返回错误，例如，
//* 应该使用错误来提高测试代码的可读性。
var AnError = errors.New("assert.AnError general error for testing")
