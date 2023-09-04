package service

import (
	"github.com/google/wire"
	"github.com/miiy/go-examples/pkg/wire/demo1/store"
)

type UserService struct {
	store *store.UserStore
}

func NewUserService(store *store.UserStore) (*UserService, error) {
	return &UserService{store: store}, nil
}

func (s *UserService) GetUserName(id int64) (string, error) {
	user, err := s.store.GetUser(id)
	if err != nil {
		return "", err
	}
	return user.Name, nil
}

var ProviderSet = wire.NewSet(store.UserStoreSet, NewUserService)
