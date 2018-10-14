package main

import (
	"log"

	"github.com/go-learn/util/connection"
	"github.com/go-learn/util/env"
	"github.com/joeshaw/envdecode"
)

type (
	config struct {
		APP      appConfig
		DATABASE databaseConfig
		REDIS    redisConfig
	}

	appConfig struct {
		NAME  string `env:"APP_NAME"`
		KEY   string `env:"APP_KEY"`
		DEBUG bool   `env:"APP_DEBUG"`
		PORT  int    `env:"APP_PORT"`
		ENV   string `env:"APP_ENV"`
	}

	databaseConfig struct {
		HOST     string `env:"DB_HOST"`
		PORT     int    `env:"DB_PORT"`
		DATABASE string `env:"DB_DATABASE"`
		USERNAME string `env:"DB_USERNAME"`
		PASSWORD string `env:"DB_PASSWORD"`
	}

	redisConfig struct {
		HOST   string `env:"REDIS_HOST"`
		PORT   int    `env:"REDIS_PORT"`
		PREFIX string `env:"REDIS_PREFIX"`
		TTL    int    `env:"REDIS_TTL"`
		WAIT   bool   `env:"REDIS_WAIT"`
	}
)

func main() {
	println("APP ENV:", env.Info())

	var cfg config
	err := envdecode.Decode(&cfg)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err)
	}

	db := connection.InitDB(
		connection.DB_CONFIG{
			HOST:     cfg.DATABASE.HOST,
			PORT:     cfg.DATABASE.PORT,
			DATABASE: cfg.DATABASE.DATABASE,
			USERNAME: cfg.DATABASE.USERNAME,
			PASSWORD: cfg.DATABASE.PASSWORD,
		},
	)

	redis := connection.InitRedis(
		connection.REDIS_CONFIG{
			HOST: cfg.REDIS.HOST,
			PORT: cfg.REDIS.PORT,
			WAIT: cfg.REDIS.WAIT,
		},
	)
}
