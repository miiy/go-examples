// +build wireinject

package main

import (
	"github.com/google/wire"
	"log"
)

func InitApplication(path string) *application {
	wire.Build(newConfig, newDB, newApplication)
	return &application{}
}

func newConfig(path string) *config {
	// read config
	dbAddr := "localhost:6379"

	return &config{
		dbAddr: dbAddr,
	}
}

func newDB(config *config) *db {
	log.Println("connection db: " + config.dbAddr)
	return &db{}
}

func newApplication(config *config, db *db) *application {
	return &application{
		config: config,
		db: db,
	}
}
