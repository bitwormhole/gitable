package ptable

import (
	"github.com/bitwormhole/starter/io/fs"
)

// DataDir 指向一个文件夹，是 Database 的容器
type DataDir interface {
	Path() fs.Path
	OpenDatabase(name string, init bool) (Database, error)
}

// DataDirFactory 是 DataDir 的生产者
type DataDirFactory interface {
	// Init(dir fs.Path) (DataDir, error)
	Open(dir fs.Path, init bool) (DataDir, error)
}
