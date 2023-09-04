//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"
import "github.com/miiy/go-examples/pkg/wire/demo1/service"

func initUserService(conf string) (*service.UserService, error) {
	panic(wire.Build(service.ProviderSet))
}
