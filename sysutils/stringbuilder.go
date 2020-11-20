package sysutils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// TStringBuilder .
type TStringBuilder struct {
	buffer *bytes.Buffer
}

// NewStringBuilder 新建String
func NewStringBuilder() *TStringBuilder {
	p := &TStringBuilder{}
	p.Clear()
	return p
}

// Append 追加
func (me *TStringBuilder) Append(s string) *TStringBuilder {
	me.buffer.WriteString(s)
	return me
}

//
// AppendStrings .
func (me *TStringBuilder) AppendStrings(ss ...string) *TStringBuilder {
	for i := range ss {
		me.buffer.WriteString(ss[i])
	}
	return me
}

// AppendInt .
func (me *TStringBuilder) AppendInt(i int) *TStringBuilder {
	me.buffer.WriteString(strconv.Itoa(i))
	return me
}

// AppendInt64 .
func (me *TStringBuilder) AppendInt64(i int64) *TStringBuilder {
	me.buffer.WriteString(strconv.FormatInt(i, 10))
	return me
}

// AppendFloat64 .
func (me *TStringBuilder) AppendFloat64(f float64) *TStringBuilder {
	me.buffer.WriteString(strconv.FormatFloat(f, 'E', -1, 64))
	return me
}

// Replace .
func (me *TStringBuilder) Replace(old, new string) *TStringBuilder {
	str := strings.Replace(me.ToString(), old, new, -1)
	me.Clear()
	me.buffer.WriteString(str)
	return me
}

// Clear .
func (me *TStringBuilder) Clear() *TStringBuilder {
	buffer := &bytes.Buffer{}
	me.buffer = buffer
	return me
}

// ToString .
func (me *TStringBuilder) ToString() string {
	return me.buffer.String()
}

// AppendFormat 格式化
func (me *TStringBuilder) AppendFormat(format string, a ...interface{}) *TStringBuilder {
	s := fmt.Sprintf(format, a...)
	me.Append(s)
	return me
}
