package model

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type UserWithPasswordAuth struct {
	User
	Password []byte `json:"-"`
}

func (UserWithPasswordAuth) TableName() string {
	return "users"
}
