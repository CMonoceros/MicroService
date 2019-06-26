package sql

import (
	"SnowBrick-Backend/common/log"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"sync/atomic"
	"time"
)

type OrmDB struct {
	*gorm.DB
	read []*gorm.DB
	idx  int64
	conf *Config
}

func OpenOrm(c *Config) (*OrmDB, error) {
	ormDB := new(OrmDB)
	ormDB.conf = c

	d, err := connectGORM(c, c.DSN)
	if err != nil {
		return nil, err
	}
	ormDB.DB = d

	rs := make([]*gorm.DB, 0, len(c.ReadDSN))
	for _, rd := range c.ReadDSN {
		d, err := connectGORM(c, rd)
		if err != nil {
			return nil, err
		}
		rs = append(rs, d)
	}
	ormDB.read = rs

	return ormDB, nil
}

func connectGORM(c *Config, dsnConfig *DSNConfig) (*gorm.DB, error) {
	d, err := gorm.Open("mysql", dsnConfig.URI)
	if err != nil {
		err = errors.WithStack(err)
		return nil, err
	}
	d.LogMode(c.LogMode)
	d.SetLogger(log.GetORMDefaultWriter(dsnConfig.Name))
	d.DB().SetMaxOpenConns(c.Active)
	d.DB().SetMaxIdleConns(c.Idle)
	d.DB().SetConnMaxLifetime(time.Duration(c.IdleTimeout))

	return d, nil
}

func (db *OrmDB) readIndex() int {
	if len(db.read) == 0 {
		return 0
	}
	v := atomic.AddInt64(&db.idx, 1)
	return int(v) % len(db.read)
}

func (db *OrmDB) ping(c context.Context, now *gorm.DB) (err error) {
	_, c, cancel := db.conf.ExecTimeout.Shrink(c)
	err = now.DB().PingContext(c)
	cancel()
	if err != nil {
		err = errors.WithStack(err)
	}
	return
}

func (db *OrmDB) Ping(c context.Context) (err error) {
	if err = db.ping(c, db.DB); err != nil {
		return
	}
	for _, rd := range db.read {
		if err = db.ping(c, rd); err != nil {
			return
		}
	}
	return
}

func (db *OrmDB) Close() (err error) {
	if e := db.DB.Close(); e != nil {
		err = errors.WithStack(e)
	}
	for _, rd := range db.read {
		if e := rd.Close(); e != nil {
			err = errors.WithStack(e)
		}
	}
	return
}

func (db *OrmDB) ReadOnlyTable(name string) *gorm.DB {
	idx := db.readIndex()
	for i := range db.read {
		if rd := db.read[(idx+i)%len(db.read)].Table(name); rd != nil {
			return rd
		}
	}
	return db.Table(name)
}

func (db *OrmDB) ReadOnlyModel(value interface{}) *gorm.DB {
	idx := db.readIndex()
	for i := range db.read {
		if rd := db.read[(idx+i)%len(db.read)].Model(value); rd != nil {
			return rd
		}
	}
	return db.Model(value)
}

func (db *OrmDB) ReadOnly() *gorm.DB {
	idx := db.readIndex()
	for i := range db.read {
		if rd := db.read[(idx+i)%len(db.read)]; rd != nil {
			return rd
		}
	}
	return db.DB
}
