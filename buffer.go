package gopcua

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"time"

	"github.com/wmnsk/gopcua/utils"
)

const (
	null    = 0xffffffff
	f32qnan = 0xffc00000
	f64qnan = 0xfff8000000000000
)

type Buffer struct {
	buf []byte
	pos int
	err error
}

func NewBuffer(b []byte) *Buffer {
	return &Buffer{buf: b}
}

func (b *Buffer) Error() error {
	return b.err
}

func (b *Buffer) Bytes() []byte {
	if b.err != nil {
		return nil
	}
	return b.buf[b.pos:]
}

func (b *Buffer) Pos() int {
	return b.pos
}

func (b *Buffer) Len() int {
	return len(b.buf) - b.pos
}

func (b *Buffer) ReadBool() bool {
	return b.ReadByte() > 0
}

func (b *Buffer) ReadByte() byte {
	if b.err != nil {
		return 0
	}
	d := b.ReadN(1)
	if d == nil {
		return 0
	}
	return d[0]
}

func (b *Buffer) ReadInt8() int8 {
	return int8(b.ReadByte())
}

func (b *Buffer) ReadInt16() int16 {
	return int16(b.ReadUint16())
}

func (b *Buffer) ReadUint16() uint16 {
	if b.err != nil {
		return 0
	}
	d := b.ReadN(2)
	if d == nil {
		return 0
	}
	return binary.LittleEndian.Uint16(d)
}

func (b *Buffer) ReadInt32() int32 {
	return int32(b.ReadUint32())
}

func (b *Buffer) ReadUint32() uint32 {
	if b.err != nil {
		return 0
	}
	d := b.ReadN(4)
	if d == nil {
		return 0
	}
	return binary.LittleEndian.Uint32(d)
}

func (b *Buffer) ReadInt64() int64 {
	return int64(b.ReadUint64())
}

func (b *Buffer) ReadUint64() uint64 {
	if b.err != nil {
		return 0
	}
	d := b.ReadN(8)
	if d == nil {
		return 0
	}
	return binary.LittleEndian.Uint64(d)
}

func (b *Buffer) ReadFloat32() float32 {
	if b.err != nil {
		return 0
	}
	bits := b.ReadUint32()
	if b.err != nil {
		return 0
	}
	return math.Float32frombits(bits)
}

func (b *Buffer) ReadFloat64() float64 {
	if b.err != nil {
		return 0
	}
	bits := b.ReadUint64()
	if b.err != nil {
		return 0
	}
	return math.Float64frombits(bits)
}

func (b *Buffer) ReadString() string {
	return string(b.ReadBytes())
}

func (b *Buffer) ReadBytes() []byte {
	n := b.ReadUint32()
	if b.err != nil {
		return nil
	}
	if n == 0 || n == null {
		return nil
	}
	d := b.ReadN(int(n))
	if b.err != nil {
		return nil
	}
	return d
}

func (b *Buffer) ReadStruct(r BinaryDecoder) {
	if b.err != nil {
		return
	}
	n, err := r.Decode(b.buf[b.pos:])
	if err != nil {
		b.err = err
		return
	}
	b.pos += n
}

func (b *Buffer) ReadTime() time.Time {
	d := b.ReadN(8)
	if b.err != nil {
		return time.Time{}
	}
	return utils.DecodeTimestamp(d)
}

func (b *Buffer) ReadN(n int) []byte {
	if b.err != nil {
		return nil
	}
	d := b.buf[b.pos:]
	if n > len(d) {
		b.err = io.ErrUnexpectedEOF
		return nil
	}
	b.pos += n
	return d[:n]
}

func (b *Buffer) WriteBool(v bool) {
	if v {
		b.WriteUint8(1)
	} else {
		b.WriteUint8(0)
	}
}

func (b *Buffer) WriteByte(n byte) {
	b.buf = append(b.buf, n)
}

func (b *Buffer) WriteInt8(n int8) {
	b.buf = append(b.buf, byte(n))
}

func (b *Buffer) WriteUint8(n uint8) {
	b.buf = append(b.buf, byte(n))
}

func (b *Buffer) WriteInt16(n int16) {
	b.WriteUint16(uint16(n))
}

func (b *Buffer) WriteUint16(n uint16) {
	d := make([]byte, 2)
	binary.LittleEndian.PutUint16(d, n)
	b.Write(d)
}

func (b *Buffer) WriteInt32(n int32) {
	b.WriteUint32(uint32(n))
}

func (b *Buffer) WriteUint32(n uint32) {
	d := make([]byte, 4)
	binary.LittleEndian.PutUint32(d, n)
	b.Write(d)
}

func (b *Buffer) WriteInt64(n int64) {
	b.WriteUint64(uint64(n))
}

func (b *Buffer) WriteUint64(n uint64) {
	d := make([]byte, 8)
	binary.LittleEndian.PutUint64(d, n)
	b.Write(d)
}

func (b *Buffer) WriteFloat32(n float32) {
	bits := math.Float32bits(n)
	b.WriteUint32(bits)
}

func (b *Buffer) WriteFloat64(n float64) {
	bits := math.Float64bits(n)
	b.WriteUint64(bits)
}

func (b *Buffer) WriteString(s string) {
	if s == "" {
		b.WriteUint32(null)
		return
	}
	b.WriteByteString([]byte(s))
}

func (b *Buffer) WriteByteString(d []byte) {
	if b.err != nil {
		return
	}
	if len(d) > math.MaxInt32 {
		b.err = fmt.Errorf("value too large")
		return
	}
	if d == nil {
		b.WriteUint32(null)
		return
	}
	b.WriteUint32(uint32(len(d)))
	b.Write(d)
}

func (b *Buffer) WriteStruct(w BinaryEncoder) {
	if b.err != nil {
		return
	}
	var d []byte
	d, b.err = w.Encode()
	b.Write(d)
}

func (b *Buffer) WriteTime(v time.Time) {
	d := make([]byte, 8)
	utils.EncodeTimestamp(d, v)
	b.Write(d)
}

func (b *Buffer) Write(d []byte) {
	if b.err != nil {
		return
	}
	b.buf = append(b.buf, d...)
}
