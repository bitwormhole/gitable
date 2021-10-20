package engine

import (
	"errors"
	"strings"

	"github.com/bitwormhole/ptable"
	"github.com/bitwormhole/starter/collection"
)

type rowImpl struct {
	session   ptable.Session
	table     ptable.Table
	rowKey    string
	tableName string
	keyPrefix string                // 'tableName.rowKey.'
	props     collection.Properties // cache for props
}

func (inst *rowImpl) _Impl() ptable.Row {
	return inst
}

func (inst *rowImpl) getProps() collection.Properties {
	p := inst.props
	if p == nil {
		p = inst.session.GetProperties(inst.table)
		inst.props = p
	}
	return p
}

func (inst *rowImpl) Owner() ptable.Table {
	return inst.table
}

func (inst *rowImpl) Key() string {
	return inst.rowKey
}

func (inst *rowImpl) Delete() error {
	prefix := inst.keyPrefix
	p := inst.getProps()
	all := p.Export(nil)
	p.Clear()
	count := 0
	for k, v := range all {
		if strings.HasPrefix(k, prefix) {
			count++
			continue
		}
		p.SetProperty(k, v)
	}
	if count <= 0 {
		return errors.New("no row with id:" + prefix)
	}
	return nil
}

func (inst *rowImpl) SetValue(field string, value string) error {
	p := inst.getProps()
	p.SetProperty(inst.keyPrefix+field, value)
	return nil
}

func (inst *rowImpl) GetValue(field string) (string, error) {
	p := inst.getProps()
	value := p.GetProperty(inst.keyPrefix+field, "")
	return value, nil
}

func (inst *rowImpl) Exists() bool {
	const field = "uuid"
	key := inst.keyPrefix + field
	p := inst.getProps()
	value := p.GetProperty(key, "")
	return len(value) > 10
}
