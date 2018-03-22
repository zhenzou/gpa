package common

import (
	"bytes"
	"io"
)

func NewOutputBuffer(buf []byte) *OutputBuffer {
	buffer := bytes.NewBuffer(nil)
	buffer.Write(buf)
	return &OutputBuffer{
		buf: buffer,
	}
}

type OutputBuffer struct {
	buf *bytes.Buffer
}

func (o *OutputBuffer) WriteRune(r rune) (n int, err error) {
	return o.buf.WriteRune(r)
}

func (o *OutputBuffer) WriteRuneAt(r rune, off int) (n int, err error) {
	if off > o.buf.Len() {
		return 0, io.EOF
	}
	bs := o.buf.Bytes()
	end := make([]byte, o.buf.Len()-off)
	copy(end, bs[off:])
	buf := bytes.NewBuffer(bs[:off])
	n, err = buf.WriteRune(r)
	buf.Write(end)
	o.buf = buf
	return
}

func (o *OutputBuffer) WriteString(str string) (n int, err error) {
	return o.buf.WriteString(str)
}

func (o *OutputBuffer) WriteStringAt(str string, off int) (n int, err error) {
	return o.WriteAt([]byte(str), off)
}

func (o *OutputBuffer) Write(p []byte) (n int, err error) {
	return o.buf.Write(p)
}

func (o *OutputBuffer) WriteAt(p []byte, off int) (n int, err error) {
	if off > o.buf.Len() {
		return 0, io.EOF
	}
	bs := o.buf.Bytes()
	end := make([]byte, o.buf.Len()-off)
	copy(end, bs[off:])
	buf := bytes.NewBuffer(bs[:off])
	n, err = buf.Write(p)
	buf.Write(end)
	o.buf = buf
	return
}

func (o *OutputBuffer) WriteByte(b byte) (err error) {
	return o.buf.WriteByte(b)
}

func (o *OutputBuffer) WriteByteAt(b byte, off int) (err error) {
	_, err = o.WriteAt([]byte{b}, off)
	return
}

func (o *OutputBuffer) String() string {
	return o.buf.String()
}

func (o *OutputBuffer) Bytes() []byte {
	return o.buf.Bytes()
}
