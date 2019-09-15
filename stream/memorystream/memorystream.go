package ms

import (
	"bytes"
	"encoding/binary"
)

// TMemoryStream 内存流
type TMemoryStream struct {
	buff *bytes.Buffer
}

// NewMemoryStream 构造
func NewMemoryStream() *TMemoryStream {
	p := &TMemoryStream{}
	p.init()
	return p
}

func (self *TMemoryStream) init() {
	self.buff = bytes.NewBuffer([]byte{})

}

// WriteBytes 写入一串二进制
func (self *TMemoryStream) WriteBytes(buff []byte) (int, error) {
	return self.buff.Write(buff)
}

//

// Write .
func (self *TMemoryStream) Write(data interface{}, order binary.ByteOrder) error {
	return binary.Write(self.buff, order, data)
}

//

// WriteBig .
func (self *TMemoryStream) WriteBig(data interface{}) error {
	return binary.Write(self.buff, binary.BigEndian, data)
}

// WriteLittle .
func (self *TMemoryStream) WriteLittle(data interface{}) error {
	return binary.Write(self.buff, binary.LittleEndian, data)
}

// Length .
func (self *TMemoryStream) Length() int {
	return self.buff.Len()
}

// Size .
func (self *TMemoryStream) Size() int {
	return self.buff.Len()
}

// Len .
func (self *TMemoryStream) Len() int {
	return self.buff.Len()
}

// ToArray .
func (self *TMemoryStream) ToArray() []byte {
	return self.buff.Bytes()
}

// ToBytes .
func (self *TMemoryStream) ToBytes() []byte {
	return self.buff.Bytes()
}

// Bytes .
func (self *TMemoryStream) Bytes() []byte {
	return self.buff.Bytes()
}

// // WriteInt 插入整数 大端
// func (self *TMemoryStream) WriteInt(value int) {
// 	binary.Write(self.buff, binary.BigEndian, value)
// }
