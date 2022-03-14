package database

import (
	"github.com/NexClipper/sudory/pkg/server/database/prepare"
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

type XormTableNameHolder interface {
	TableName() string
}

// XormContext
type XormContext struct {
	tx *xorm.Session
}

func NewXormContext(session *xorm.Session) Context {
	return &XormContext{tx: session}
}

// Close
//  close session
func (context XormContext) Close() error {
	return context.tx.Close()
}

// Prepared
func (context XormContext) Prepared(preparer prepare.Preparer) Context {
	context.tx = preparer.Prepared(context.tx)
	return context
}

// Where
func (context XormContext) Where(where string, args ...interface{}) Context {
	context.tx = context.tx.Where(where, args...)
	return context
}

// Create
func (context XormContext) Create(record interface{}) error {
	affect, err := context.tx.Insert(record)
	if err != nil {
		return errors.Wrapf(err, "xorm insert record=%#v", record)
	} else if !(0 < affect) {
		return errors.Wrapf(ErrorNoAffected(), "xorm insert record=%#v", record)
	}

	return nil
}

// Get
func (context XormContext) Get(record interface{}) error {
	if has, err := context.tx.Get(record); err != nil {
		return errors.Wrapf(err, "xorm get record=%#v", record)
	} else if !has {
		return errors.Wrapf(ErrorRecordWasNotFound(), "xorm get record=%#v", record)
	}

	return nil
}

// Find
func (context XormContext) Find(records interface{}) error {
	if err := context.tx.Find(records); err != nil {
		return errors.Wrapf(err, "xorm find records=%#v", records)
	}

	return nil
}

// Update
func (context XormContext) Update(record interface{}) error {
	if affect, err := context.tx.Update(record); err != nil {
		return errors.Wrapf(err, "xorm update record=%#v", record)
	} else if !(0 < affect) {
		return errors.Wrapf(ErrorNoAffected(), "xorm update record=%#v", record)
	}

	return nil
}

// Delete
func (context XormContext) Delete(record interface{}) error {
	if affect, err := context.tx.Delete(record); err != nil {
		return errors.Wrapf(err, "xorm delete record=%#v", record)
	} else if !(0 < affect) {
		return nil //idempotent
	}

	return nil
}