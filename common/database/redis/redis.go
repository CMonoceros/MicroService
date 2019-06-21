package redis

import (
	"SnowBrick-Backend/common/ctime"
	"github.com/gomodule/redigo/redis"
	"time"
)

type PoolConfig struct {
	Active      int
	Idle        int
	IdleTimeout ctime.Duration
	WaitTimeout ctime.Duration
	Wait        bool
}

type Config struct {
	*PoolConfig

	Proto       string
	DB          int
	Addr        string
	Auth        string
	DialTimeout ctime.Duration
}

func NewPool(cfg *Config) *redis.Pool {
	return &redis.Pool{
		MaxIdle:         cfg.PoolConfig.Idle,
		IdleTimeout:     time.Duration(cfg.PoolConfig.IdleTimeout),
		MaxActive:       cfg.PoolConfig.Active,
		Wait:            cfg.PoolConfig.Wait,
		MaxConnLifetime: time.Duration(cfg.DialTimeout),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(cfg.Proto, cfg.Addr)
			if err != nil {
				return nil, err
			}

			if cfg.Auth != "" {
				if _, err := c.Do("AUTH", cfg.Auth); err != nil {
					c.Close()
					return nil, err
				}
			}

			if cfg.DB != 0 {
				if _, err := c.Do("SELECT", cfg.DB); err != nil {
					c.Close()
					return nil, err
				}
				return c, nil
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Duration(cfg.PoolConfig.WaitTimeout) {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
