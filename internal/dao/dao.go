package dao

import (
	"context"

	redigo "github.com/gomodule/redigo/redis"

	"SnowBrick-Backend/common/database/redis"
	"SnowBrick-Backend/common/database/sql"
	"SnowBrick-Backend/common/log"
	"SnowBrick-Backend/conf"
)

// Dao dao.
type Dao struct {
	db    *sql.OrmDB
	redis *redigo.Pool
}

// New new a dao and return.
func New(c *conf.Config) (dao *Dao) {
	dao = &Dao{
		// mysql
		db: sql.NewMySQL(c.Mysql),
		// redis
		redis: redis.NewPool(c.Redis),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	if d.db != nil {
		d.db.Close()
	}
	if d.redis != nil {
		d.redis.Close()
	}
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.pingRedis(); err != nil {
		return
	}
	return d.db.Ping(ctx)
}

func (d *Dao) pingRedis() (err error) {
	conn := d.redis.Get()
	defer conn.Close()
	if _, err = conn.Do("PING"); err != nil {
		log.Error("conn.PING error(%v)", err)
	}
	return
}
