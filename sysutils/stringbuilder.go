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
func (m *TStringBuilder) Append(s string) *TStringBuilder {
	m.buffer.WriteString(s)
	return m
}

//
// AppendStrings .
func (m *TStringBuilder) AppendStrings(ss ...string) *TStringBuilder {
	for i := range ss {
		m.buffer.WriteString(ss[i])
	}
	return m
}

// AppendInt .
func (m *TStringBuilder) AppendInt(i int) *TStringBuilder {
	m.buffer.WriteString(strconv.Itoa(i))
	return m
}

// AppendInt64 .
func (m *TStringBuilder) AppendInt64(i int64) *TStringBuilder {
	m.buffer.WriteString(strconv.FormatInt(i, 10))
	return m
}

// AppendFloat64 .
func (m *TStringBuilder) AppendFloat64(f float64) *TStringBuilder {
	m.buffer.WriteString(strconv.FormatFloat(f, 'E', -1, 64))
	return m
}

// Replace .
func (m *TStringBuilder) Replace(old, new string) *TStringBuilder {
	str := strings.Replace(m.ToString(), old, new, -1)
	m.Clear()
	m.buffer.WriteString(str)
	return m
}

// Clear .
func (m *TStringBuilder) Clear() *TStringBuilder {
	buffer := &bytes.Buffer{}
	m.buffer = buffer
	return m
}

// ToString .
func (m *TStringBuilder) ToString() string {
	return m.buffer.String()
}

// AppendFormat 格式化
func (m *TStringBuilder) AppendFormat(format string, a ...interface{}) *TStringBuilder {
	s := fmt.Sprintf(format, a...)
	m.Append(s)
	return m
}
