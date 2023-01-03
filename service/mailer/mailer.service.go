package service

import "github.com/adibSetiawann/cronjob/model"

type MailerService interface {
	SendEmail(user *model.SendOtp) error
	VerifyEmail(user *model.VerifyEmail) error
	Validation(customerRequest interface{}) error
}