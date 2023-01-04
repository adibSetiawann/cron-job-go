package service

import "github.com/adibSetiawann/cronjob/model"

type UserService interface {
	Signin(formLogin *model.LoginForm) (string, error)
	Signup(customerRequest model.CreateUser) (model.UserResponse, error)
	GetAllData() ([]model.UserResponse, error)
	GetById(id string) ([]model.UserResponse, error)
	Validation(customerRequest interface{}) error
	UpdateEmail(form model.LoginForm) error
}