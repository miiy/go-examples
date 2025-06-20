package store

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/miiy/go-examples/wire/demo1/config"
	"github.com/miiy/go-examples/wire/demo1/db"
)

type UserStore struct {
	cfg *config.Config
	db  *sql.DB
}

type User struct {
	Id   int64
	Name string
}

func NewUserStore(cfg *config.Config, db *sql.DB) (*UserStore, error) {
	return &UserStore{
		cfg: cfg,
		db:  db,
	}, nil
}

func (s *UserStore) GetUser(id int64) (*User, error) {
	return &User{
		Id:   1,
		Name: "zhangsan",
	}, nil
}

var UserStoreSet = wire.NewSet(NewUserStore, config.NewDefaultConfig, db.NewDB)
