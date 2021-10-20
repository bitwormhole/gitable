package engine

import "github.com/bitwormhole/ptable"

type baseColumn struct {
	tableName  string
	columnName string
	table      ptable.Table
}

func (inst *baseColumn) _Impl() ptable.Column {
	return inst
}

func (inst *baseColumn) Owner() ptable.Table {
	return inst.table
}

func (inst *baseColumn) TableName() string {
	return inst.tableName
}

func (inst *baseColumn) ColumnName() string {
	return inst.columnName
}
