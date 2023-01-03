package service

import (
	"errors"
	"log"

	"github.com/adibSetiawann/cronjob/entity"
	"github.com/adibSetiawann/cronjob/model"
	"github.com/adibSetiawann/cronjob/repository/user"
	"github.com/adibSetiawann/cronjob/utils"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func (cs *UserServiceImpl) Signin(formLogin *model.LoginForm) (string, error) {
	token, err := cs.userRepo.Signin(formLogin)
	return token, err
}

func (cs *UserServiceImpl) Signup(request model.CreateUser) (model.UserResponse, error) {
	var userResponse model.UserResponse

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		log.Println(err)
		return userResponse, nil
	}
	user := entity.User{
		Email:    request.Email,
		Password: hashedPassword,
		Status: "pending",
	}

	errCreate := cs.userRepo.Signup(&user)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return userResponse, errCreate
	}

	userResponse.ID = user.ID
	userResponse.Email = user.Email
	userResponse.Status = user.Status
	return userResponse, nil
}

func (ms *UserServiceImpl) GetById(id string) ([]model.UserResponse, error) {
	merchants, errFind := ms.userRepo.FindById(id)
	if errFind != nil {
		return nil, errFind
	}
	return merchants, nil
}

func (cs *UserServiceImpl) GetAllData() ([]model.UserResponse, error) {
	users, errFind := cs.userRepo.FindAll()
	if errFind != nil {
		return nil, errFind
	}
	return users, nil
}

func (ms *UserServiceImpl) Validation(userRequest interface{}) error {
	var messageError string
	var isError bool

	// VALIDATION USER INPUT
	validate := validator.New()
	errValidate := validate.Struct(userRequest)
	if errValidate != nil {
		messageError += errValidate.Error()
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}

func NewUserService(userRepo *repository.UserRepository) UserService {
	return &UserServiceImpl{
		userRepo: *userRepo,
	}
}
