package connection

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB_CONFIG struct {
	HOST                    string
	PORT                    int
	DATABASE                string
	USERNAME                string
	PASSWORD                string
	MAX_CONNECTION_OPEN     int
	MAX_CONNECTION_LIFETIME time.Duration
	MAX_CONNECTION_IDLE     int
}

var sqlxOpen = sqlx.Open
var logFatalf = log.Fatalf

func InitDB(cfg DB_CONFIG) *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cfg.USERNAME, cfg.USERNAME, cfg.HOST, cfg.PORT, cfg.DATABASE)
	db, err := sqlxOpen("mysql", dsn)
	if err != nil {
		logFatalf("Failed connecting to db. dsn: %s, err: %s", dsn, err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		logFatalf("Failed pinging to db. dsn: %s, err: %s", dsn, err)
		return nil
	}

	db.SetMaxOpenConns(cfg.MAX_CONNECTION_OPEN)
	db.SetMaxIdleConns(cfg.MAX_CONNECTION_IDLE)
	db.SetConnMaxLifetime(cfg.MAX_CONNECTION_LIFETIME)

	return db
}
