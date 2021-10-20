package engine

import (
	"errors"

	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/io/fs"
)

////////////////////////////////////////////////////////////////////////////////

type tableImpl struct {
	context   *Context
	file      fs.Path
	tableName string
	pkey      string
	owner     ptable.Database
}

func (inst *tableImpl) _Impl() ptable.Table {
	return inst
}

func (inst *tableImpl) PrimaryKey() string {
	return inst.pkey
}

func (inst *tableImpl) Owner() ptable.Database {
	return inst.owner
}

func (inst *tableImpl) Path() fs.Path {
	return inst.file
}

func (inst *tableImpl) Name() string {
	return inst.tableName
}

func (inst *tableImpl) initColumn(name string, col *baseColumn) {
	col.table = inst
	col.tableName = inst.tableName
	col.columnName = name
}

///////////////

func (inst *tableImpl) GetColumnString(name string) ptable.ColumnString {
	col := &myStringColumn{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnBool(name string) ptable.ColumnBool {
	col := &myBoolColumn{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnTime(name string) ptable.ColumnTime {
	col := &myTimeColumn{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

///////////////

func (inst *tableImpl) GetColumnFloat32(name string) ptable.ColumnFloat32 {
	col := &myFloat32Column{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnFloat64(name string) ptable.ColumnFloat64 {
	col := &myFloat64Column{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

///////////////

func (inst *tableImpl) GetColumnInt(name string) ptable.ColumnInt {
	col := &myColumnInt{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnInt8(name string) ptable.ColumnInt8 {
	col := &myColumnInt8{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnInt16(name string) ptable.ColumnInt16 {
	col := &myColumnInt16{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnInt32(name string) ptable.ColumnInt32 {
	col := &myColumnInt32{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnInt64(name string) ptable.ColumnInt64 {
	col := &myColumnInt64{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

///////////////

func (inst *tableImpl) GetColumnUint(name string) ptable.ColumnUint {
	col := &myColumnUInt{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnUint8(name string) ptable.ColumnUint8 {
	col := &myColumnUInt8{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnUint16(name string) ptable.ColumnUint16 {
	col := &myColumnUInt16{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnUint32(name string) ptable.ColumnUint32 {
	col := &myColumnUInt32{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

func (inst *tableImpl) GetColumnUint64(name string) ptable.ColumnUint64 {
	col := &myColumnUInt64{}
	inst.initColumn(name, &col.baseColumn)
	return col
}

////////////////////////////////////////////////////////////////////////////////

type tableFactory struct {
	context *Context
}

func (inst *tableFactory) _Impl() ptable.TableFactory {
	return inst
}

func (inst *tableFactory) Open(p *ptable.TableOpen) (ptable.Table, error) {

	if p == nil {
		return nil, errors.New("no param 'TableOpen'")
	}

	if p.PrimaryKey == "" {
		p.PrimaryKey = "id"
	}

	table := &tableImpl{}
	table.context = inst.context
	table.tableName = p.TableName
	table.file = p.File
	table.owner = p.OwnerDB
	table.pkey = p.PrimaryKey

	if p.DoInit {
		if !p.File.Exists() {
			p.File.GetIO().WriteText("", nil, true)
		}
	}

	if !p.File.IsFile() {
		return nil, errors.New("no file, path=" + p.File.Path())
	}

	return table, nil
}
