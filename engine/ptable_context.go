package engine

import "github.com/bitwormhole/ptable"

// Context 是 ptable 数据引擎上下文
type Context struct {
	DataDirFactory  ptable.DataDirFactory
	DatabaseFactory ptable.DatabaseFactory
	TableFactory    ptable.TableFactory
	SessionFactory  ptable.SessionFactory
	// SessionManager  ptable.SessionManager
}
