// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

import (
	"encoding/binary"
	"github.com/go-faster/errors"
	"math"
)

// ClickHouse uses LittleEndian.
var _ = binary.LittleEndian

// ColFloat64 represents Float64 column.
type ColFloat64 []float64

// Compile-time assertions for ColFloat64.
var (
	_ ColInput  = ColFloat64{}
	_ ColResult = (*ColFloat64)(nil)
	_ Column    = (*ColFloat64)(nil)
)

// Type returns ColumnType of Float64.
func (ColFloat64) Type() ColumnType {
	return ColumnTypeFloat64
}

// Rows returns count of rows in column.
func (c ColFloat64) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColFloat64) Reset() {
	*c = (*c)[:0]
}

// NewArrFloat64 returns new Array(Float64).
func NewArrFloat64() *ColArr {
	return &ColArr{
		Data: new(ColFloat64),
	}
}

// AppendFloat64 appends slice of float64 to Array(Float64).
func (c *ColArr) AppendFloat64(data []float64) {
	d := c.Data.(*ColFloat64)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}

// EncodeColumn encodes Float64 rows to *Buffer.
func (c ColFloat64) EncodeColumn(b *Buffer) {
	const size = 64 / 8
	offset := len(b.Buf)
	b.Buf = append(b.Buf, make([]byte, size*len(c))...)
	for _, v := range c {
		binary.LittleEndian.PutUint64(
			b.Buf[offset:offset+size],
			math.Float64bits(v),
		)
		offset += size
	}
}

// DecodeColumn decodes Float64 rows from *Reader.
func (c *ColFloat64) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	const size = 64 / 8
	data, err := r.ReadRaw(rows * size)
	if err != nil {
		return errors.Wrap(err, "read")
	}
	v := append(*c, make([]float64, rows)...)

	var (
		n = 0
		i = 0
	)
	const (
		unroll         = 4
		unrollByteSize = size * unroll
	)
	if len(data) > unrollByteSize {
		_ = data[:len(data)-unrollByteSize]
		for i = 0; i <= len(data)-unrollByteSize; i += unrollByteSize {
			src := [unroll]float64{}
			const offset0 = 0 + size + size + size
			src[3] =
				math.Float64frombits(bin.Uint64(data[i+offset0 : i+offset0+size : i+offset0+size]))
			const offset1 = 0 + size + size
			src[2] =
				math.Float64frombits(bin.Uint64(data[i+offset1 : i+offset1+size : i+offset1+size]))
			const offset2 = 0 + size
			src[1] =
				math.Float64frombits(bin.Uint64(data[i+offset2 : i+offset2+size : i+offset2+size]))
			const offset3 = 0
			src[0] =
				math.Float64frombits(bin.Uint64(data[i+offset3 : i+offset3+size : i+offset3+size]))
			copy(v[n:n+unroll:n+unroll], src[:])
			n += unroll
		}
	}
	for _ = i; i < len(data); i += size {
		v[n] =
			math.Float64frombits(bin.Uint64(data[i:]))
		n++
	}
	*c = v
	return nil
}
