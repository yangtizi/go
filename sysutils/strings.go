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

// TrimEx 去掉指定的内容
func TrimEx(s string, tm string) string {
	return strings.Trim(s, tm)
}

// TrimLeft 去掉左边
func TrimLeft(s string) string {
	return strings.TrimLeft(s, " ")
}

// TrimRight 去掉右边
func TrimRight(s string) string {
	return strings.TrimRight(s, " ")
}

// ToLower 转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper 转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Compare 比较 0： a==b,  -1 ： a < b,   +1 ： a > b.
func Compare(a, b string) int {
	return strings.Compare(a, b)
}

// IndexOf 字符串匹配
// sysutils.IndexOf("slisa", "is") // 2
func IndexOf(s, sub string) int {
	return strings.Index(s, sub) // -1; 没找到
}

// LastIndexOf 字符串匹配
// sysutils.LastIndexOf("slisa", "is") // 2
func LastIndexOf(s, sub string) int {
	return strings.LastIndex(s, sub)
}

// IndexOfAny 字符串匹配，后面是任意包含
// sysutils.IndexOfAny("slisa", "is") // 0
func IndexOfAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}

// LastIndexOfAny 字符串匹配，后面是任意包含
// sysutils.LastIndexOfAny("slisa", "is") // 3
func LastIndexOfAny(s, chars string) int {
	return strings.LastIndexAny(s, chars)
}

// Contains 是否包含
func Contains(s, sub string) bool {
	return strings.Contains(s, sub)
}

// StartsWith 判断是否以某字符串打头
func StartsWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// EndsWith 判断是否以某字符串打结尾
func EndsWith(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// QuotedString 补充冒号， 通常用来做显示用
func QuotedString(s string) string {
	return strconv.Quote(s)
}

// DeQuotedString 去掉冒号， 通常用来做显示用
func DeQuotedString(s string) (string, error) {
	return strconv.Unquote(s)
}

// PadLeft 左边补齐长度
func PadLeft(s string, totalWidth int) string {
	strPad := ""
	for i := 0; i < totalWidth-len(s); i++ {
		strPad += " "
	}
	return strPad + s
}

// PadRight 右边补齐长度
func PadRight(s string, totalWidth int) string {
	strPad := ""
	for i := 0; i < totalWidth-len(s); i++ {
		strPad += " "
	}

	return s + strPad
}

// Substring 字符串截取
func Substring(s string, nStartIndex, nLength int) string {
	rs := []rune(s)
	nRuneLen := len(rs)
	nEndIndex := 0

	// 从后往前的
	if nStartIndex < 0 {
		nStartIndex = nRuneLen - 1 + nStartIndex
	}
	nEndIndex = nStartIndex + nLength

	if nStartIndex > nEndIndex {
		nStartIndex, nEndIndex = nEndIndex, nStartIndex
	}

	if nStartIndex < 0 {
		nStartIndex = 0
	} else if nStartIndex > nRuneLen {
		nStartIndex = nRuneLen
	}
	if nEndIndex < 0 {
		nEndIndex = 0
	} else if nEndIndex > nRuneLen {
		nEndIndex = nRuneLen
	}

	return string(rs[nStartIndex:nEndIndex])
}

// Split 分割
// usage strArr := sysutils.Split("A1,B2,C3,D4,E5,F6,G7", ",")
func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

// Join 字符串组合
// usage sysutils.Join(strArr, "-") // A1-B2-C3-D4-E5-F6-G7
func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}
