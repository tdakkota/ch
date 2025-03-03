// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColFixedStr512 represents FixedStr512 column.
type ColFixedStr512 [][512]byte

// Compile-time assertions for ColFixedStr512.
var (
	_ ColInput  = ColFixedStr512{}
	_ ColResult = (*ColFixedStr512)(nil)
	_ Column    = (*ColFixedStr512)(nil)
)

// Rows returns count of rows in column.
func (c ColFixedStr512) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColFixedStr512) Reset() {
	*c = (*c)[:0]
}

// Type returns ColumnType of FixedStr512.
func (ColFixedStr512) Type() ColumnType {
	return ColumnTypeFixedString.With("512")
}

// Row returns i-th row of column.
func (c ColFixedStr512) Row(i int) [512]byte {
	return c[i]
}

// Append [512]byte to column.
func (c *ColFixedStr512) Append(v [512]byte) {
	*c = append(*c, v)
}

// Append [512]byte slice to column.
func (c *ColFixedStr512) AppendArr(vs [][512]byte) {
	*c = append(*c, vs...)
}

// LowCardinality returns LowCardinality for FixedStr512.
func (c *ColFixedStr512) LowCardinality() *ColLowCardinality[[512]byte] {
	return &ColLowCardinality[[512]byte]{
		index: c,
	}
}

// Array is helper that creates Array of [512]byte.
func (c *ColFixedStr512) Array() *ColArr[[512]byte] {
	return &ColArr[[512]byte]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable([512]byte).
func (c *ColFixedStr512) Nullable() *ColNullable[[512]byte] {
	return &ColNullable[[512]byte]{
		Values: c,
	}
}

// NewArrFixedStr512 returns new Array(FixedStr512).
func NewArrFixedStr512() *ColArr[[512]byte] {
	return &ColArr[[512]byte]{
		Data: new(ColFixedStr512),
	}
}
