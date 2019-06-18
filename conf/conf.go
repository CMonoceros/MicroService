package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/log"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"

	xtime "github.com/bilibili/kratos/pkg/time"
)

type Config struct {
	Env           string
	TimeoutSecond int
	Log           *log.Config
	BM            *bm.ServerConfig
	Warden        *warden.ServerConfig
	Mysql         *sql.Config
	Redis         *redis.Config
	RedisExpire   xtime.Duration
}

var (
	confPath string
	Conf     = &Config{}
)

func init() {
	flag.StringVar(&confPath, "conf", "", "config path")
}

func Init() error {
	if confPath != "" {
		return local()
	}
	return nil
}

func local() (err error) {
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}
