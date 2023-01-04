package repository

import (
	"time"

	"github.com/adibSetiawann/cronjob/config"
	"github.com/adibSetiawann/cronjob/entity"
	"github.com/adibSetiawann/cronjob/model"
	"github.com/adibSetiawann/cronjob/utils"
	"github.com/golang-jwt/jwt"
)

type UserRepositoryImplement struct {
}

func (*UserRepositoryImplement) Signin(loginForm *model.LoginForm) (string, error) {
	var userData entity.User

	err := config.DB.Debug().First(&userData, "email=?", loginForm.Email)
	if err.Error != nil {
		return "user not found in database", err.Error
	}
	isValid, errValid := utils.ConfirmPassword(loginForm.Password, userData.Password)

	if !isValid {
		return "please input correct password", errValid
	}

	claims := jwt.MapClaims{}
	claims["email"] = userData.Email
	claims["exp"] = time.Now().Add(time.Minute * 1500).Unix()
	if userData.Email == "admin@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		return "failed generate token", errToken
	}

	return token, nil
}

func (*UserRepositoryImplement) Signup(user *entity.User) error {
	
	db := config.DB.Debug().Create(&user)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (*UserRepositoryImplement) UpdateEmail(user *model.LoginForm) error {
	var userData entity.User

	err := config.DB.Debug().First(&userData, "email=?", user.Email)
	if err.Error != nil {
		return err.Error
	}

	userData.Email = user.Email
	userData.Status = "pending"

	db := config.DB.Debug().Save(&userData)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (*UserRepositoryImplement) FindById(id string) ([]model.UserResponse, error) {

	var users []model.UserResponse

	err := config.DB.Debug().Preload("Wallets").Preload("Mailers").Preload("Currency").First(&users, "id=?", id)
	if err.Error != nil {
		return nil, err.Error
	}

	return users, nil
}

func (*UserRepositoryImplement) FindAll() ([]model.UserResponse, error) {

	var users []model.UserResponse

	db := config.DB.Debug().Preload("Wallets").Preload("Mailers").Find(&users)
	if db.Error != nil {
		return nil, db.Error
	}

	return users, nil
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImplement{}
}
