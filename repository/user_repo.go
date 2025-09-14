package repository

import "github.com/amos-babu/auth-jwt/model"

type UserRepo interface {
	Register(user *model.User) (*model.User, error)
	Login() (*model.User, error)
	Logout() error
}
