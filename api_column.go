package ptable

import "time"

// Column 负责某个字段（列）的访问
type Column interface {
	Owner() Table
	TableName() string
	ColumnName() string
}

// ColumnX ...
type ColumnX interface {
	Column
}

// ColumnString ...
type ColumnString interface {
	Column
	Get(row Row) string
	Set(row Row, value string)
}

// ColumnFloat32 ...
type ColumnFloat32 interface {
	Column
	Get(row Row) float32
	Set(row Row, value float32)
}

// ColumnFloat64 ...
type ColumnFloat64 interface {
	Column
	Get(row Row) float64
	Set(row Row, value float64)
}

// ColumnInt ...
type ColumnInt interface {
	Column
	Get(row Row) int
	Set(row Row, value int)
}

// ColumnInt8 ...
type ColumnInt8 interface {
	Column
	Get(row Row) int8
	Set(row Row, value int8)
}

// ColumnInt16 ...
type ColumnInt16 interface {
	Column
	Get(row Row) int16
	Set(row Row, value int16)
}

// ColumnInt32 ...
type ColumnInt32 interface {
	Column
	Get(row Row) int32
	Set(row Row, value int32)
}

// ColumnInt64 ...
type ColumnInt64 interface {
	Column
	Get(row Row) int64
	Set(row Row, value int64)
}

// ColumnUint ...
type ColumnUint interface {
	Column
	Get(row Row) uint
	Set(row Row, value uint)
}

// ColumnUint8 ...
type ColumnUint8 interface {
	Column
	Get(row Row) uint8
	Set(row Row, value uint8)
}

// ColumnUint16 ...
type ColumnUint16 interface {
	Column
	Get(row Row) uint16
	Set(row Row, value uint16)
}

// ColumnUint32 ...
type ColumnUint32 interface {
	Column
	Get(row Row) uint32
	Set(row Row, value uint32)
}

// ColumnUint64 ...
type ColumnUint64 interface {
	Column
	Get(row Row) uint64
	Set(row Row, value uint64)
}

// ColumnBool ...
type ColumnBool interface {
	Column
	Get(row Row) bool
	Set(row Row, value bool)
}

// ColumnTime ...
type ColumnTime interface {
	Column
	Get(row Row) time.Time
	Set(row Row, value time.Time)
}
