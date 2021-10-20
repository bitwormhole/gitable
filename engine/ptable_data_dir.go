package engine

import (
	"errors"
	"strings"

	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/io/fs"
)

////////////////////////////////////////////////////////////////////////////////

type dataDirFactoryImpl struct {
	context *Context
}

func (inst *dataDirFactoryImpl) _Impl() ptable.DataDirFactory {
	return inst
}

func (inst *dataDirFactoryImpl) makeInstance(dir fs.Path) *dataDirImpl {
	dd := &dataDirImpl{}
	dd.context = inst.context
	dd.dir = dir
	dd.dotPTable = dir.GetChild(".ptable")
	return dd
}

func (inst *dataDirFactoryImpl) Init(dir fs.Path) (ptable.DataDir, error) {
	dd := inst.makeInstance(dir)
	if dd.dotPTable.Exists() {
		return nil, errors.New("the data-dir is exists, path=" + dir.Path())
	}
	err := dd.dir.Mkdirs()
	if err != nil {
		return nil, err
	}
	err = dd.dotPTable.GetIO().WriteText("ptable.version=1\n", nil, true)
	if err != nil {
		return nil, err
	}
	return dd, nil
}

func (inst *dataDirFactoryImpl) Open(dir fs.Path, init bool) (ptable.DataDir, error) {
	dd := inst.makeInstance(dir)
	if dd.dotPTable.IsFile() && dd.dir.IsDir() {
		return dd, nil
	}
	if init {
		return inst.Init(dir)
	}
	return nil, errors.New("cannot open data-dir at path=" + dir.Path())
}

////////////////////////////////////////////////////////////////////////////////

type dataDirImpl struct {
	context   *Context
	dir       fs.Path
	dotPTable fs.Path // file of '.ptable'
}

func (inst *dataDirImpl) _Impl() ptable.DataDir {
	return inst
}

func (inst *dataDirImpl) Path() fs.Path {
	return inst.dir
}

func (inst *dataDirImpl) OpenDatabase(name string, init bool) (ptable.Database, error) {

	name = strings.TrimSpace(name)
	if len(name) < 1 {
		return nil, errors.New("bad db name: " + name)
	}

	dir := inst.dir.GetChild(name)
	return inst.context.DatabaseFactory.Open(name, dir, inst, init)
}
