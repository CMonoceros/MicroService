package sql

import (
	"SnowBrick-Backend/common/ctime"
	"SnowBrick-Backend/common/log"

	// database driver
	_ "github.com/go-sql-driver/mysql"
)

// Config mysql dsn config.
type DSNConfig struct {
	URI  string
	Name string
}

// Config mysql config.
type Config struct {
	LogMode      bool           // for log mode
	Addr         string         // for trace
	DSN          *DSNConfig     // write data source name.
	ReadDSN      []*DSNConfig   // read data source name.
	Active       int            // pool
	Idle         int            // pool
	IdleTimeout  ctime.Duration // connect max life time.
	QueryTimeout ctime.Duration // query sql timeout
	ExecTimeout  ctime.Duration // execute sql timeout
	TranTimeout  ctime.Duration // transaction sql timeout
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *OrmDB) {
	if c.QueryTimeout == 0 || c.ExecTimeout == 0 || c.TranTimeout == 0 {
		panic("mysql must be set query/execute/transction timeout")
	}
	db, err := OpenOrm(c)
	if err != nil {
		log.Error("open mysql error(%v)", err)
		panic(err)
	}
	return
}
