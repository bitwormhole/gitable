package engine

import (
	"strconv"

	"github.com/bitwormhole/ptable"
)

const (
	ibits   = 0
	i8bits  = 8
	i16bits = 16
	i32bits = 32
	i64bits = 64
)

////////////////////////////////////////////////////////////////////////////////

type myColumnInt struct {
	baseColumn
}

func (inst *myColumnInt) _Impl() ptable.ColumnInt {
	return inst
}

func (inst *myColumnInt) Get(row ptable.Row) int {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseInt(text, 10, ibits)
	if err != nil {
		return 0
	}
	return int(n)
}

func (inst *myColumnInt) Set(row ptable.Row, value int) {
	text := strconv.FormatInt(int64(value), 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myColumnInt8 struct {
	baseColumn
}

func (inst *myColumnInt8) _Impl() ptable.ColumnInt8 {
	return inst
}

func (inst *myColumnInt8) Get(row ptable.Row) int8 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseInt(text, 10, i8bits)
	if err != nil {
		return 0
	}
	return int8(n)
}

func (inst *myColumnInt8) Set(row ptable.Row, value int8) {
	text := strconv.FormatInt(int64(value), 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myColumnInt16 struct {
	baseColumn
}

func (inst *myColumnInt16) _Impl() ptable.ColumnInt16 {
	return inst
}

func (inst *myColumnInt16) Get(row ptable.Row) int16 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseInt(text, 10, i16bits)
	if err != nil {
		return 0
	}
	return int16(n)
}

func (inst *myColumnInt16) Set(row ptable.Row, value int16) {
	text := strconv.FormatInt(int64(value), 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myColumnInt32 struct {
	baseColumn
}

func (inst *myColumnInt32) _Impl() ptable.ColumnInt32 {
	return inst
}

func (inst *myColumnInt32) Get(row ptable.Row) int32 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseInt(text, 10, i32bits)
	if err != nil {
		return 0
	}
	return int32(n)
}

func (inst *myColumnInt32) Set(row ptable.Row, value int32) {
	text := strconv.FormatInt(int64(value), 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myColumnInt64 struct {
	baseColumn
}

func (inst *myColumnInt64) _Impl() ptable.ColumnInt64 {
	return inst
}

func (inst *myColumnInt64) Get(row ptable.Row) int64 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseInt(text, 10, i64bits)
	if err != nil {
		return 0
	}
	return n
}

func (inst *myColumnInt64) Set(row ptable.Row, value int64) {
	text := strconv.FormatInt(value, 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////
