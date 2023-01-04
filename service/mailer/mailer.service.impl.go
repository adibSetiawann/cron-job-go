package service

import (
	"errors"

	"github.com/adibSetiawann/cronjob/model"
	repo "github.com/adibSetiawann/cronjob/repository/mailer"
	"github.com/go-playground/validator/v10"
)

type MailerServiceImpl struct {
	mailerRepo repo.MailerRepository
}

func (ms *MailerServiceImpl) SendEmail(payload *model.SendOtp) error{

	data := model.SendOtp{
		Email: payload.Email,
		UserId: payload.UserId,
	}

	_, err := ms.mailerRepo.SendOtp(&data)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MailerServiceImpl) VerifyEmail(payload *model.VerifyEmail) error{

	err := ms.mailerRepo.VerifiedEmail(payload.Email)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MailerServiceImpl) Validation(payload interface{}) error {
	var messageError string
	var isError bool

	// VALIDATION USER INPUT
	validate := validator.New()
	errValidate := validate.Struct(payload)
	if errValidate != nil {
		messageError += errValidate.Error()
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}

func NewMailerService(mailerRepo *repo.MailerRepository) MailerService {
	return &MailerServiceImpl{
		mailerRepo: *mailerRepo,
	}
}
