package ptable

import (
	"io"

	"github.com/bitwormhole/starter/collection"
)

// Row 指向具体的一行记录
type Row interface {
	Owner() Table
	Key() string
	Delete() error
	SetValue(field string, value string) error
	GetValue(field string) (string, error)
	Exists() bool
}

// Session 会话对象
type Session interface {
	io.Closer
	BeginTransaction() Transaction
	DB() Database
	GetProperties(table Table) collection.Properties
	ListIDs(table Table) []string
	GetRow(table Table, key string) Row
	GetRowRequired(table Table, key string) (Row, error)
}

// SessionFactory 会话工厂
type SessionFactory interface {
	OpenSession(db Database) (Session, error)
}

// SessionManager 会话管理器
type SessionManager interface {
	GetSession(db Database, create bool) (Session, error)
}

// Transaction 事务对象
type Transaction interface {
	io.Closer
	Commit()
	Rollback()
}
