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
func (self *TStringBuilder) Append(s string) *TStringBuilder {
	self.buffer.WriteString(s)
	return self
}

//
// AppendStrings .
func (self *TStringBuilder) AppendStrings(ss ...string) *TStringBuilder {
	for i := range ss {
		self.buffer.WriteString(ss[i])
	}
	return self
}

// AppendInt .
func (self *TStringBuilder) AppendInt(i int) *TStringBuilder {
	self.buffer.WriteString(strconv.Itoa(i))
	return self
}

// AppendInt64 .
func (self *TStringBuilder) AppendInt64(i int64) *TStringBuilder {
	self.buffer.WriteString(strconv.FormatInt(i, 10))
	return self
}

// AppendFloat64 .
func (self *TStringBuilder) AppendFloat64(f float64) *TStringBuilder {
	self.buffer.WriteString(strconv.FormatFloat(f, 'E', -1, 64))
	return self
}

// Replace .
func (self *TStringBuilder) Replace(old, new string) *TStringBuilder {
	str := strings.Replace(self.ToString(), old, new, -1)
	self.Clear()
	self.buffer.WriteString(str)
	return self
}

// Clear .
func (self *TStringBuilder) Clear() *TStringBuilder {
	buffer := &bytes.Buffer{}
	self.buffer = buffer
	return self
}

// ToString .
func (self *TStringBuilder) ToString() string {
	return self.buffer.String()
}

// AppendFormat 格式化
func (self *TStringBuilder) AppendFormat(format string, a ...interface{}) *TStringBuilder {
	s := fmt.Sprintf(format, a...)
	self.Append(s)
	return self
}
