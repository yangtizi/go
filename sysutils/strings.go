package sysutils

import (
	"encoding/hex"
	"math/big"
	"strconv"
	"strings"
	"time"
)

// 小数
type float interface {
	float32 | float64
}

// 整数
type integer interface {
	int32 | int64 | uint32 | uint64 | int
}

// 实数
type number interface {
	float | integer
	// func String() string;
}

// 取绝对值
func Abs[V number](n V) V {
	if n < 0 {
		return -n
	}
	return n
}

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

// StrToUint 字符串转无符号整数
func StrToUint(str string) (uint, error) {
	n, err := StrToUint64(str)
	return uint(n), err
}

// StrToUintDef 字符串转整数 带默认值
func StrToUintDef(str string, nDef uint) uint {
	n, err := StrToUint(str)

	if err != nil {
		return nDef
	}
	return n
}

// StrToUint64 字符串转无符号64整数
func StrToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 0)
}

// StrToIntDef 字符串转整数 带默认值
func StrToUint64Def(str string, nDef uint64) uint64 {
	n, err := StrToUint64(str)

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

// IntToHex 整数转成HEX字符串
func IntToHex(n int) string {
	return strconv.FormatInt(int64(n), 16)
}

// Int64ToHex 整数转成HEX字符串
func Int64ToHex(n int64) string {
	return strconv.FormatInt(n, 16)
}

// Uint64ToHex 整数转成HEX字符串
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

// BigIntToStr 大数转字符串
func BigIntToStr(b *big.Int) string {
	return b.Text(10)
}

// BigIntToHex 大数转HEX
func BigIntToHex(b *big.Int) string {
	return b.Text(16)
}

// FloatToStr 整数转字符串
// sysutils.FloatToStr(1.20, 2)   // 1.2
// sysutils.FloatToStr(1.2222, 2) // 1.22
// sysutils.FloatToStr(1.2022, 2) // 1.2
// sysutils.FloatToStr(1.000, 2)  // 1
func FloatToStr(f float64, nPrec int) string {
	strFloat := strconv.FormatFloat(f, 'f', nPrec, 64)

	nIndex := LastIndexOf(strFloat, ".")
	if nIndex <= -1 {

	} else {
		strFloat = strings.TrimRight(strFloat, "0") // 去掉后面的0
		strFloat = strings.TrimRight(strFloat, ".") // 去掉后面的.
	}

	return strFloat
}

// Float64ToStr
func Float64ToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func StrToFloat(s string) (float64, error) {
	return StrToFloat64(s)
}

// StrToFloat64
func StrToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func StrToFloatDef(s string, fDef float64) float64 {
	return StrToFloat64Def(s, fDef)
}

func StrToFloat64Def(s string, fDef float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fDef
	}
	return f
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
// sysutils.IndexOf("slisa", "is") // sl[is]a 2
// sysutils.IndexOf("slisa", "s") // [s]lisa 0
func IndexOf(s, sub string) int {
	return strings.Index(s, sub) // -1; 没找到
}

// LastIndexOf 字符串匹配
// sysutils.LastIndexOf("slisa", "is") // sl[is]a 2
// sysutils.LastIndexOf("slisa", "s") // sli[s]a 3
func LastIndexOf(s, sub string) int {
	return strings.LastIndex(s, sub)
}

// IndexOfAny 字符串匹配，后面是任意包含
// sysutils.IndexOfAny("slisa", "is") // [s]lisa 0
// sysutils.IndexOfAny("slisa", "as") // [s]lisa 0
// sysutils.IndexOfAny("slisa", "li") // s[l]isa 1
func IndexOfAny(s, chars string) int {
	return strings.IndexAny(s, chars)
}

// LastIndexOfAny 字符串匹配，后面是任意包含
// sysutils.LastIndexOfAny("slisa", "as") // slis[a] 4
// sysutils.LastIndexOfAny("slisa", "is") // sli[s]a 3
// sysutils.LastIndexOfAny("slisa", "li") // sl[i]sa 2
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

// DateToStr 日期转字符串
func DateToStr(datetime time.Time) string {
	return datetime.Format("2006-01-02")
}

// FormatDateTimeStr 切换传统字符串变成golang字符串
// 返回值是一种格式化后的字符串,重点来看Format参数中的指令字符
// FormatDateTimeStr('c',now);
// 输出为：2004-8-7 10:26:58
// FormatDateTimeStr('yy-mm-dd',now);
// FormatDateTimeStr('yy\mm\dd',now);
// 输出为： 04-08-07
// 也可以用":"来分开时间
// FormatDateTimeStr('hh:nn:ss',now);
// 输出为：10:32:23
func FormatDateTimeStr(strOrigin string) string {
	switch strOrigin {
	case "c":
		// c 以短时间格式显示时间，即全部是数字的表示
		// 输出为：2004-8-7 9:55:40
		return "2006-01-02 15:04:05"
	case "ddddd":
		// ddddd 以短时间格式显示年月日
		// 输出为：2004-8-7
		return "2006-01-02"
	case "dddddd":
		// dddddd 以长时间格式显示年月日
		// 输出为：2004年8月7日
		return "2006年1月2日"
	case "t":
		// t 以短时间格式显示时间
		// 输出为 10:17
		return "15:04"
	case "tt":
		// tt 以长时间格式显示时间
		// 输出为10:18:46
		return "15:04:05"

	}

	// d 对应于时间中的日期，日期是一位则显示一位，两位则显示两位
	// 输出可能为1～31
	// dd 和d的意义一样，但它始终是以两位来显示的
	// 输出可能为01～31
	strOrigin = StringReplaceAll(strOrigin, "dd", "02")
	strOrigin = StringReplaceAll(strOrigin, "d", "2")

	// h/hh,n/nn,s/ss,z/zzz 分别表示小时，分，秒,毫秒
	strOrigin = StringReplaceAll(strOrigin, "zzz", ".000")
	strOrigin = StringReplaceAll(strOrigin, "z", ".999")

	strOrigin = StringReplaceAll(strOrigin, "ss", "05")
	strOrigin = StringReplaceAll(strOrigin, "s", "5")

	strOrigin = StringReplaceAll(strOrigin, "nn", "04")
	strOrigin = StringReplaceAll(strOrigin, "n", "4")

	strOrigin = StringReplaceAll(strOrigin, "hh", "15")
	strOrigin = StringReplaceAll(strOrigin, "h", ".3")

	// yy/yyyy 表示年
	// FormatdateTime('yy',now);
	// 输出为 04
	// FormatdateTime('yyyy',now);
	// 输出为 2004,
	strOrigin = StringReplaceAll(strOrigin, "yyyy", "2006")
	strOrigin = StringReplaceAll(strOrigin, "yy", "06")

	// FormatdateTime('m',now);
	// 输出为：8
	// FormatdateTime('mm',now);
	// 输出为 08
	strOrigin = StringReplaceAll(strOrigin, "mm", "01")
	strOrigin = StringReplaceAll(strOrigin, "m", "1")

	return strOrigin
}
