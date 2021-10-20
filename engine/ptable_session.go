package engine

import (
	"errors"
	"sort"
	"strings"

	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/buffers"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

////////////////////////////////////////////////////////////////////////////////

type sessionFactoryImpl struct{}

func (inst *sessionFactoryImpl) _Impl() ptable.SessionFactory {
	return inst
}

func (inst *sessionFactoryImpl) OpenSession(db ptable.Database) (ptable.Session, error) {
	session := &sessionImpl{}
	session.db = db
	return session._Impl(), nil
}

////////////////////////////////////////////////////////////////////////////////

type sessionManagerImpl struct {
	context *Context
	table   map[string]ptable.Session
}

func (inst *sessionManagerImpl) _Impl() ptable.SessionManager {
	return inst
}

func (inst *sessionManagerImpl) getTable() map[string]ptable.Session {
	t := inst.table
	if t == nil {
		t = make(map[string]ptable.Session)
		inst.table = t
	}
	return t
}

func (inst *sessionManagerImpl) GetSession(db ptable.Database, create bool) (ptable.Session, error) {
	table := inst.getTable()
	name := db.Name()
	older := table[name]
	if older != nil {
		return older, nil
	}
	if create {
		factory := inst.context.SessionFactory
		pNew, err := factory.OpenSession(db)
		if err != nil {
			return nil, err
		}
		older = pNew
		table[name] = pNew
		return pNew, nil
	}
	return nil, errors.New("no session for database: " + name)
}

////////////////////////////////////////////////////////////////////////////////

type sessionImpl struct {
	db       ptable.Database
	cache    map[string]*tableCache
	trans    *transactionImpl
	revision int64
}

func (inst *sessionImpl) _Impl() ptable.Session {
	return inst
}

func (inst *sessionImpl) nextRevision() int64 {
	inst.revision++
	return inst.revision
}

func (inst *sessionImpl) getCache() map[string]*tableCache {
	cacheT := inst.cache
	if cacheT == nil {
		cacheT = make(map[string]*tableCache)
		inst.cache = cacheT
	}
	return cacheT
}

func (inst *sessionImpl) Close() error {
	// inst.db = nil
	inst.cache = nil
	inst.trans = nil
	return nil
}

func (inst *sessionImpl) DB() ptable.Database {
	return inst.db
}

func (inst *sessionImpl) BeginTransaction() ptable.Transaction {

	tr := inst.trans
	if tr != nil {
		if tr.isClosed() {
			tr = nil
		}
	}

	if tr != nil {
		return tr
	}

	tr = &transactionImpl{}
	tr.init(inst)

	inst.trans = tr
	return tr
}

func (inst *sessionImpl) GetProperties(table ptable.Table) collection.Properties {
	cacheT := inst.getCache()
	c := cacheT[table.Name()]
	if c == nil {
		c = &tableCache{}
		c.init(inst, table)
		cacheT[table.Name()] = c
	}
	return c.getProps()
}

func (inst *sessionImpl) GetRow(table ptable.Table, key string) (ptable.Row, error) {
	row := &rowImpl{}
	row.table = table
	row.session = inst
	row.tableName = table.Name()
	row.rowKey = key
	row.keyPrefix = row.tableName + "." + row.rowKey + "."
	return row, nil
}

func (inst *sessionImpl) ListIDs(table ptable.Table) []string {

	prefix := table.Name() + "."
	suffix := "." + table.PrimaryKey()
	props := inst.GetProperties(table)
	all := props.Export(nil)
	ids := make([]string, 0)

	for key := range all {
		if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
			id := key[len(prefix) : len(key)-len(suffix)]
			ids = append(ids, id)
		}
	}

	return ids
}

////////////////////////////////////////////////////////////////////////////////

type tableCache struct {
	session *sessionImpl
	table   ptable.Table
	buffer  buffers.TextFileBuffer
	file    fs.Path

	revision int64
	content  string
	props    collection.Properties
}

func (inst *tableCache) init(session *sessionImpl, table ptable.Table) {

	file := table.Path()

	inst.session = session
	inst.buffer.Init(file)
	inst.file = file
	inst.table = table

	inst.revision = 0
}

func (inst *tableCache) getProps() collection.Properties {

	have := inst.content
	want := inst.buffer.GetText(false)
	props := inst.props

	if have != want {
		props = nil
		have = want
		inst.content = want
	}

	if props != nil {
		return props
	}

	// reload
	pInner, err := collection.ParseProperties(want, nil)
	if err != nil {
		vlog.Warn(err)
		pInner = collection.CreateProperties()
	}
	props = (&sessionPropertiesWrapper{}).init(inst, pInner)

	// keep
	inst.props = props
	return props
}

func (inst *tableCache) commit() {
	props := inst.props
	if props == nil {
		return
	}
	text := inst.stringifyProperties(props)
	inst.buffer.SetText(text, false)
}

func (inst *tableCache) stringifyProperties(props collection.Properties) string {
	if props == nil {
		return ""
	}
	builder := strings.Builder{}
	list := make([]string, 0)
	all := props.Export(nil)
	for k, v := range all {
		list = append(list, k+" = "+v)
	}
	sort.Strings(list)
	for _, line := range list {
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////

type transactionImpl struct {
	session  *sessionImpl
	revision int64
	closed   bool
}

func (inst *transactionImpl) _Impl() ptable.Transaction {
	return inst
}

func (inst *transactionImpl) init(session *sessionImpl) ptable.Transaction {
	inst.session = session
	inst.revision = session.nextRevision()
	return inst
}

func (inst *transactionImpl) isClosed() bool {
	return inst.closed
}

func (inst *transactionImpl) Close() error {
	inst.closed = true
	return nil
}

func (inst *transactionImpl) listModifiedTables() []*tableCache {
	from := inst.revision
	to := inst.session.nextRevision()
	tables := inst.session.getCache()
	results := make([]*tableCache, 0)
	for _, tc := range tables {
		if tc == nil {
			continue
		}
		if from <= tc.revision && tc.revision <= to {
			results = append(results, tc)
		}
	}
	return results
}

func (inst *transactionImpl) Commit() {

	if inst.closed {
		return
	}
	inst.closed = true

	// 提交所有修改
	list := inst.listModifiedTables()
	for _, tc := range list {
		tc.commit()
	}
}

func (inst *transactionImpl) Rollback() {

	if inst.closed {
		return
	}
	inst.closed = true

	// 放弃所有修改
	list := inst.listModifiedTables()
	for _, tc := range list {
		tc.props = nil
	}
}

////////////////////////////////////////////////////////////////////////////////

type sessionPropertiesWrapper struct {
	session *sessionImpl
	cache   *tableCache
	inner   collection.Properties
}

func (inst *sessionPropertiesWrapper) init(cache *tableCache, inner collection.Properties) collection.Properties {
	inst.cache = cache
	inst.inner = inner
	inst.session = cache.session
	return inst
}

func (inst *sessionPropertiesWrapper) update() {
	inst.cache.revision = inst.session.nextRevision()
}

func (inst *sessionPropertiesWrapper) GetPropertyRequired(name string) (string, error) {
	return inst.inner.GetPropertyRequired(name)
}

func (inst *sessionPropertiesWrapper) GetProperty(name string, defaultValue string) string {
	return inst.inner.GetProperty(name, defaultValue)
}

func (inst *sessionPropertiesWrapper) SetProperty(name string, value string) {
	inst.inner.SetProperty(name, value)
	inst.update()
}

func (inst *sessionPropertiesWrapper) Clear() {
	inst.inner.Clear()
	inst.update()
}

func (inst *sessionPropertiesWrapper) Export(dst map[string]string) map[string]string {
	return inst.inner.Export(dst)
}

func (inst *sessionPropertiesWrapper) Import(src map[string]string) {
	inst.inner.Import(src)
	inst.update()
}

////////////////////////////////////////////////////////////////////////////////
