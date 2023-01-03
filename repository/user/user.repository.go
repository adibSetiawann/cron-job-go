package repository

import (
	"github.com/adibSetiawann/cronjob/entity"
	"github.com/adibSetiawann/cronjob/model"
)

type UserRepository interface {
	Signin(form *model.LoginForm) (string, error)
	Signup(customer *entity.User) error
	FindAll() ([]model.UserResponse, error)
	FindById(id string) ([]model.UserResponse, error)
}
