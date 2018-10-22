package user

import (
	"context"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
)

type core struct {
	db    *sqlx.DB
	redis *redis.Pool
}

type ICore interface {
	GetUserListFromDB() ([]TodoModel, error)
}

var logFatalf = log.Fatalf

func Init(db *sqlx.DB, redis *redis.Pool) ICore {
	return &core{
		db:    db,
		redis: redis,
	}
}

func examineDBHealth(db *sqlx.DB) {
	if db == nil {
		logFatalf("Failed to initialize user. db object cannot be nil")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := db.PingContext(ctx)
	if err != nil {
		logFatalf("Failed to initialize user. cannot pinging to db. err: %s", err)
		return
	}
}

func examineRedisHealth(redis *redis.Pool) {
	if redis == nil {
		logFatalf("Failed to initialize user. redis object cannot be nil")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := redis.GetContext(ctx)
	if err != nil {
		logFatalf("Failed to initialize user. cannot connect to redis. err: %s", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("PING")
	if err != nil {
		logFatalf("Failed to initialize user. cannot pinging to redis. err: %s", err)
		return
	}
}
