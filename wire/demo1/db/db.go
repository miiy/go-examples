package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/miiy/go-examples/wire/demo1/config"
)

// NewDB return a db instance
// https://github.com/go-sql-driver/mysql
func NewDB(c *config.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", c.DSN)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, err
}
