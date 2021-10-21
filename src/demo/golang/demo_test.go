package golang

import (
	"errors"
	"strconv"
	"testing"

	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/ptable/engine"
	"github.com/bitwormhole/starter/io/fs"
)

func TestDemo(t *testing.T) {
	err := doTestDemo(t)
	if err != nil {
		t.Error(err)
	}
}

func doTestDemo(t *testing.T) error {

	dir := fs.Default().GetPath(t.TempDir()).GetChild("data")
	factory := engine.DefaultFactory()

	dd, err := factory.Open(dir, true)
	if err != nil {
		return err
	}

	db, err := dd.OpenDatabase("demo-db-1", true)
	if err != nil {
		return err
	}

	session, err := db.OpenSession()
	if err != nil {
		return err
	}

	tr := session.BeginTransaction()

	repo := &DemoRepoImpl{db: db, session: session}
	dao := repo.init()

	entity := &DemoEntity{
		id: 1,
		f1: 2,
		f2: "3",
		f3: true,
	}
	entity, err = dao.Insert(entity)
	if err != nil {
		return err
	}

	tr.Commit()

	list := dao.All()
	for x := range list {
		t.Log(x)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

type DemoEntity struct {
	id int

	f1 int
	f2 string
	f3 bool
}

type DemoRepo interface {
	Insert(e *DemoEntity) (*DemoEntity, error)
	Delete(id string) error
	Update(id string, e *DemoEntity) (*DemoEntity, error)
	Find(id string) (*DemoEntity, error)
	All() []*DemoEntity
	GetIDs() []string
}

////////////////////////////////////////////////////////////////////////////////

type DemoRepoImpl struct {
	db      ptable.Database
	table   ptable.Table
	session ptable.Session

	pk ptable.ColumnInt
	f1 ptable.ColumnInt
	f2 ptable.ColumnString
	f3 ptable.ColumnBool
}

func (inst *DemoRepoImpl) init() DemoRepo {

	tableOpen := &ptable.TableOpen{}
	tableOpen.TableName = "table1demo"
	tableOpen.DoInit = true
	tableOpen.PrimaryKey = "id"

	db := inst.db
	table, _ := db.OpenTable(tableOpen)

	inst.pk = table.GetColumnInt("id")
	inst.f1 = table.GetColumnInt("f1")
	inst.f2 = table.GetColumnString("f2")
	inst.f3 = table.GetColumnBool("f3")

	inst.table = table
	return inst
}

func (inst *DemoRepoImpl) load(row ptable.Row, e *DemoEntity) {
	e.id = inst.pk.Get(row)
	e.f1 = inst.f1.Get(row)
	e.f2 = inst.f2.Get(row)
	e.f3 = inst.f3.Get(row)
}

func (inst *DemoRepoImpl) save(row ptable.Row, e *DemoEntity) {
	inst.pk.Set(row, e.id)
	inst.f1.Set(row, e.f1)
	inst.f2.Set(row, e.f2)
	inst.f3.Set(row, e.f3)
}

func (inst *DemoRepoImpl) Insert(e *DemoEntity) (*DemoEntity, error) {
	session := inst.session
	table := inst.table
	row := session.GetRow(table, strconv.Itoa(e.id))
	if row.Exists() {
		return nil, errors.New("row is exists")
	}
	inst.save(row, e)
	return e, nil
}

func (inst *DemoRepoImpl) Delete(id string) error {
	session := inst.session
	table := inst.table
	row, err := session.GetRowRequired(table, id)
	if err != nil {
		return err
	}
	err = row.Delete()
	if err != nil {
		return err
	}
	return nil
}

func (inst *DemoRepoImpl) Update(id string, e *DemoEntity) (*DemoEntity, error) {
	session := inst.session
	table := inst.table
	row, err := session.GetRowRequired(table, strconv.Itoa(e.id))
	if err != nil {
		return nil, err
	}
	inst.save(row, e)
	return e, nil
}

func (inst *DemoRepoImpl) Find(id string) (*DemoEntity, error) {
	session := inst.session
	table := inst.table
	row, err := session.GetRowRequired(table, id)
	if err != nil {
		return nil, err
	}
	e := &DemoEntity{}
	inst.load(row, e)
	return e, nil
}

func (inst *DemoRepoImpl) All() []*DemoEntity {
	session := inst.session
	table := inst.table
	ids := session.ListIDs(table)
	list := make([]*DemoEntity, 0)
	for _, id := range ids {
		e := &DemoEntity{}
		row, err := session.GetRowRequired(table, id)
		if err != nil {
			break
		}
		inst.load(row, e)
		list = append(list, e)
	}
	return list
}

func (inst *DemoRepoImpl) GetIDs() []string {
	session := inst.session
	table := inst.table
	return session.ListIDs(table)
}
