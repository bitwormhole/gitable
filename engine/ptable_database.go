package engine

import (
	"errors"
	"strings"

	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/io/fs"
)

////////////////////////////////////////////////////////////////////////////////

type databaseImpl struct {
	context *Context
	name    string
	dir     fs.Path
	owner   ptable.DataDir
}

func (inst *databaseImpl) _Impl() ptable.Database {
	return inst
}

func (inst *databaseImpl) Name() string {
	return inst.name
}

func (inst *databaseImpl) Path() fs.Path {
	return inst.dir
}

func (inst *databaseImpl) Owner() ptable.DataDir {
	return inst.owner
}

func (inst *databaseImpl) OpenTable(p *ptable.TableOpen) (ptable.Table, error) {
	if p == nil {
		return nil, errors.New("no param")
	}
	p = p.Clone()
	name := strings.TrimSpace(p.TableName)
	if len(name) < 1 {
		return nil, errors.New("bad table name: " + name)
	}
	file := inst.dir.GetChild(name + ".ptable")
	p.File = file
	p.TableName = name
	return inst.context.TableFactory.Open(p)
}

func (inst *databaseImpl) OpenSession() (ptable.Session, error) {
	factory := inst.context.SessionFactory
	return factory.OpenSession(inst)
}

////////////////////////////////////////////////////////////////////////////////

type databaseFactory struct {
	context *Context
}

func (inst *databaseFactory) _Impl() ptable.DatabaseFactory {
	return inst
}

func (inst *databaseFactory) Open(name string, dir fs.Path, owner ptable.DataDir, init bool) (ptable.Database, error) {

	db := &databaseImpl{}
	db.context = inst.context
	db.name = name
	db.dir = dir
	db.owner = owner

	if init {
		if !dir.Exists() {
			dir.Mkdirs()
		}
	}

	if !dir.IsDir() {
		return nil, errors.New("no dir, path=" + dir.Path())
	}

	return db, nil
}
