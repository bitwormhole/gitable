package engine

import (
	"strconv"
	"time"

	"github.com/bitwormhole/ptable"
)

////////////////////////////////////////////////////////////////////////////////

type myStringColumn struct {
	baseColumn
}

func (inst *myStringColumn) _Impl() ptable.ColumnString {
	return inst
}

func (inst *myStringColumn) Get(row ptable.Row) string {
	value, err := row.GetValue(inst.columnName)
	if err != nil {
		return ""
	}
	return value
}

func (inst *myStringColumn) Set(row ptable.Row, value string) {
	row.SetValue(inst.columnName, value)
}

////////////////////////////////////////////////////////////////////////////////

type myBoolColumn struct {
	baseColumn
}

func (inst *myBoolColumn) _Impl() ptable.ColumnBool {
	return inst
}

func (inst *myBoolColumn) Get(row ptable.Row) bool {
	value, err := row.GetValue(inst.columnName)
	if err != nil {
		return false
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return b
}

func (inst *myBoolColumn) Set(row ptable.Row, value bool) {
	str := strconv.FormatBool(value)
	row.SetValue(inst.columnName, str)
}

////////////////////////////////////////////////////////////////////////////////

type myTimeColumn struct {
	baseColumn
}

func (inst *myTimeColumn) _Impl() ptable.ColumnTime {
	return inst
}

func (inst *myTimeColumn) Get(row ptable.Row) time.Time {
	// todo
	return time.Unix(0, 0)
}

func (inst *myTimeColumn) Set(row ptable.Row, value time.Time) {
	// todo
	row.SetValue(inst.columnName, "0")
}

////////////////////////////////////////////////////////////////////////////////
