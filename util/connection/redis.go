package connection

import (
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

type REDIS_CONFIG struct {
	HOST       string
	PORT       int
	MAX_ACTIVE int
	WAIT       bool
}

type RedisConn struct {
	Core *redis.Pool
}

var redisDial = redis.Dial

func InitRedis(cfg REDIS_CONFIG) (conn *redis.Pool) {
	var address = cfg.HOST + ":" + strconv.Itoa(cfg.PORT)
	return &redis.Pool{
		IdleTimeout:     2 * time.Second,
		MaxConnLifetime: 10 * time.Second,
		MaxActive:       cfg.MAX_ACTIVE,
		Wait:            cfg.WAIT,
		Dial: func() (redis.Conn, error) {
			return redisDial("tcp", address)
		},
	}
}
