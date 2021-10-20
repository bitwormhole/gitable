package engine

import (
	"strconv"

	"github.com/bitwormhole/ptable"
)

const (
	f32bits = 32
	f64bits = 64
)

////////////////////////////////////////////////////////////////////////////////

type myFloat32Column struct {
	baseColumn
}

func (inst *myFloat32Column) _Impl() ptable.ColumnFloat32 {
	return inst
}

func (inst *myFloat32Column) Get(row ptable.Row) float32 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseFloat(text, f32bits)
	if err != nil {
		return 0
	}
	return float32(n)
}

func (inst *myFloat32Column) Set(row ptable.Row, value float32) {
	n := float64(value)
	text := strconv.FormatFloat(n, 'f', -1, f32bits)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////

type myFloat64Column struct {
	baseColumn
}

func (inst *myFloat64Column) _Impl() ptable.ColumnFloat64 {
	return inst
}

func (inst *myFloat64Column) Get(row ptable.Row) float64 {
	text, err := row.GetValue(inst.columnName)
	if err != nil {
		return 0
	}
	n, err := strconv.ParseFloat(text, f64bits)
	if err != nil {
		return 0
	}
	return n
}

func (inst *myFloat64Column) Set(row ptable.Row, value float64) {
	text := strconv.FormatFloat(value, 'f', -1, f64bits)
	row.SetValue(inst.columnName, text)
}

////////////////////////////////////////////////////////////////////////////////
