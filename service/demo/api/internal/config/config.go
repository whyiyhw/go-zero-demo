package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	PGSql struct {
		DataSource string
	}

	RedisCache cache.CacheConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Email struct {
		Host string
		Port int
		User string
		Pass string
	}
}
