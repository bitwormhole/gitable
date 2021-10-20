package ptable

import (
	"github.com/bitwormhole/starter/io/fs"
)

// Table 指向一个文件，是 Entity 的容器
type Table interface {
	Owner() Database
	Path() fs.Path
	Name() string
	PrimaryKey() string

	GetColumnString(name string) ColumnString
	GetColumnBool(name string) ColumnBool
	GetColumnTime(name string) ColumnTime

	GetColumnFloat32(name string) ColumnFloat32
	GetColumnFloat64(name string) ColumnFloat64

	GetColumnInt(name string) ColumnInt
	GetColumnInt8(name string) ColumnInt8
	GetColumnInt16(name string) ColumnInt16
	GetColumnInt32(name string) ColumnInt32
	GetColumnInt64(name string) ColumnInt64

	GetColumnUint(name string) ColumnUint
	GetColumnUint8(name string) ColumnUint8
	GetColumnUint16(name string) ColumnUint16
	GetColumnUint32(name string) ColumnUint32
	GetColumnUint64(name string) ColumnUint64
}

// TableOpen 是打开 Table 的参数
type TableOpen struct {
	TableName  string
	PrimaryKey string // default is 'id'
	File       fs.Path
	OwnerDB    Database
	DoInit     bool
}

// TableFactory 是 Table 的生产者
type TableFactory interface {
	Open(p *TableOpen) (Table, error)
}

////////////////////////////////////////////////////////////////////////////////

func (inst *TableOpen) Clone() *TableOpen {
	n := &TableOpen{}
	n.DoInit = inst.DoInit
	n.File = inst.File
	n.OwnerDB = inst.OwnerDB
	n.PrimaryKey = inst.PrimaryKey
	n.TableName = inst.TableName
	return n
}
