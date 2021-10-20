package engine

import "github.com/bitwormhole/ptable"

// DefaultFactory 函数返回一个默认的数据引擎工厂
func DefaultFactory() ptable.DataDirFactory {

	ctx := &Context{}

	ctx.DataDirFactory = &dataDirFactoryImpl{context: ctx}
	ctx.DatabaseFactory = &databaseFactory{context: ctx}
	ctx.TableFactory = &tableFactory{context: ctx}
	ctx.SessionFactory = &sessionFactoryImpl{}

	return ctx.DataDirFactory
}
