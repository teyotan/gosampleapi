package mysqlstore

import (
	"gosampleapi/model"
)

func (s *MySQLStore) CreateUser(username string, password []byte) (*model.User, error) {
	user := model.UserWithPasswordAuth{
		User:     model.User{Username: username},
		Password: password,
	}

	err := s.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func (s *MySQLStore) GetUser(id uint64) (*model.User, error) {
	var user = model.User{ID: id}

	err := s.db.First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MySQLStore) GetUserWithPasswordByUsername(username string) (*model.UserWithPasswordAuth, error) {
	var user = model.UserWithPasswordAuth{User: model.User{Username: username}}

	err := s.db.Where(model.User{Username: username}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
