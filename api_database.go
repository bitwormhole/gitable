package ptable

import (
	"github.com/bitwormhole/starter/io/fs"
)

// Database 指向一个文件夹，是 Table 的容器
type Database interface {
	Name() string
	Path() fs.Path
	Owner() DataDir
	OpenTable(p *TableOpen) (Table, error)
	OpenSession() (Session, error)
}

// DatabaseFactory 是 Database 的生产者
type DatabaseFactory interface {
	Open(name string, dir fs.Path, owner DataDir, init bool) (Database, error)
}
