package store

import "gosampleapi/model"

type Store interface {
	GetUser(id uint64) (*model.User, error)
	GetUserWithPasswordByUsername(username string) (*model.UserWithPasswordAuth, error)
	CreateUser(username string, password []byte) (*model.User, error)
}
