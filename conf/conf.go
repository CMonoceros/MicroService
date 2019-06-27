package conf

import (
	"SnowBrick-Backend/common/oss"
	"flag"
	"github.com/BurntSushi/toml"

	"SnowBrick-Backend/common/database/redis"
	"SnowBrick-Backend/common/database/sql"
	"SnowBrick-Backend/common/log"
)

type RPCConfig struct {
	Addr    string
	Timeout string
}

type HTTPConfig struct {
	Addr    string
	Timeout string
}

type Config struct {
	Env           string
	TimeoutSecond int
	Log           *log.Config
	HTTP          *HTTPConfig
	Grpc          *RPCConfig
	Mysql         *sql.Config
	Redis         *redis.Config
	Oss           *oss.Config
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
