package engine

import (
	"strconv"

	"github.com/bitwormhole/ptable"
)

const (
	uibits   = 0
	ui8bits  = 8
	ui16bits = 16
	ui32bits = 32
	ui64bits = 64
)

////////////////////////////////////////////////////////////////////////////////

type myColumnUInt struct {
	baseColumn
}

func (inst *myColumnUInt) _Impl() ptable.ColumnUint {
	return inst
}

func (inst *myColumnUInt) Get(row ptable.Row) uint {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseUint(text, 10, uibits)
	if err != nil {
		return 0
	}
	return uint(n)
}

func (inst *myColumnUInt) Set(row ptable.Row, value uint) {
	text := strconv.FormatUint(uint64(value), 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myColumnUInt8 struct {
	baseColumn
}

func (inst *myColumnUInt8) _Impl() ptable.ColumnUint8 {
	return inst
}

func (inst *myColumnUInt8) Get(row ptable.Row) uint8 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseUint(text, 10, ui8bits)
	if err != nil {
		return 0
	}
	return uint8(n)
}

func (inst *myColumnUInt8) Set(row ptable.Row, value uint8) {
	text := strconv.FormatUint(uint64(value), 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myColumnUInt16 struct {
	baseColumn
}

func (inst *myColumnUInt16) _Impl() ptable.ColumnUint16 {
	return inst
}

func (inst *myColumnUInt16) Get(row ptable.Row) uint16 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseUint(text, 10, ui16bits)
	if err != nil {
		return 0
	}
	return uint16(n)
}

func (inst *myColumnUInt16) Set(row ptable.Row, value uint16) {
	text := strconv.FormatUint(uint64(value), 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myColumnUInt32 struct {
	baseColumn
}

func (inst *myColumnUInt32) _Impl() ptable.ColumnUint32 {
	return inst
}

func (inst *myColumnUInt32) Get(row ptable.Row) uint32 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseUint(text, 10, ui32bits)
	if err != nil {
		return 0
	}
	return uint32(n)
}

func (inst *myColumnUInt32) Set(row ptable.Row, value uint32) {
	text := strconv.FormatUint(uint64(value), 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myColumnUInt64 struct {
	baseColumn
}

func (inst *myColumnUInt64) _Impl() ptable.ColumnUint64 {
	return inst
}

func (inst *myColumnUInt64) Get(row ptable.Row) uint64 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseUint(text, 10, ui64bits)
	if err != nil {
		return 0
	}
	return n
}

func (inst *myColumnUInt64) Set(row ptable.Row, value uint64) {
	text := strconv.FormatUint(value, 10)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////
