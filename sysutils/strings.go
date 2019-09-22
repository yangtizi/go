package sysutils

import (
	"encoding/hex"
	"strconv"
	"strings"
)

// StrToInt 字符串转整数
func StrToInt(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// StrToIntDef 字符串转整数 带默认值
func StrToIntDef(str string, nDef int) int {
	n, err := strconv.Atoi(str)

	if err != nil {
		return nDef
	}
	return n
}

// Str2Int 字符串转整数
func Str2Int(str string) (int, error) {
	n, err := strconv.Atoi(str)
	return n, err
}

// Str2IntDef 字符串转整数 带默认值
func Str2IntDef(str string, nDef int) int {
	n, err := strconv.Atoi(str)

	if err != nil {
		return nDef
	}
	return n
}

// IntToStr 整数转字符串
func IntToStr(n int) string {
	return strconv.Itoa(n)
}

// Int64ToStr 整数转字符串
func Int64ToStr(n int64) string {
	return strconv.FormatInt(n, 10)
}

// Uint64ToStr 整数转字符串
func Uint64ToStr(n uint64) string {
	return strconv.FormatUint(n, 10)
}

// Uint64ToHex 整数转字符串
func Uint64ToHex(n uint64) string {
	return strconv.FormatUint(n, 16)
}

// BinToHex 二进制转HEX
func BinToHex(b []byte) string {
	return hex.EncodeToString(b)
}

// HexToBin HEX 转 二进制
func HexToBin(h string) ([]byte, error) {
	b, err := hex.DecodeString(h)
	return b, err
}

// StringReplace 替换字符串
func StringReplace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

// StringReplaceAll 替换所有字符串
func StringReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// UpperCase 转大写
func UpperCase(s string) string {
	return strings.ToUpper(s)
}

// Trim 去掉格式化
func Trim(s string) string {
	return strings.Trim(s, " ")
}
