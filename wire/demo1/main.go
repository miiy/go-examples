package main

import (
	"fmt"

	configpkg "github.com/miiy/go-examples/wire/demo1/config"
	dbpkg "github.com/miiy/go-examples/wire/demo1/db"
	"github.com/miiy/go-examples/wire/demo1/service"
	"github.com/miiy/go-examples/wire/demo1/store"
)

func main() {
	_ = withOutWire()
	_ = useWire()
}

func withOutWire() error {
	config := configpkg.NewDefaultConfig("./config.yaml")

	db, err := dbpkg.NewDB(config)
	if err != nil {
		return err
	}

	userStore, err := store.NewUserStore(config, db)
	if err != nil {
		return err
	}

	userService, err := service.NewUserService(userStore)
	if err != nil {
		return err
	}

	username, err := userService.GetUserName(1)
	fmt.Println("username:", username)
	return nil
}

func useWire() error {
	userService, err := initUserService("./config.yaml")
	if err != nil {
		return err
	}
	username, err := userService.GetUserName(1)
	fmt.Println("username:", username)
	return nil
}
