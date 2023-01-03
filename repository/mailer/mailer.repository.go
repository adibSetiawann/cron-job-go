package repository

import (
	"github.com/adibSetiawann/cronjob/model"
)

type MailerRepository interface {
	VerifyEmail(customer *model.VerifyEmail) error
	SendOtp(customer *model.SendOtp) (string, error)
	ExpireLink(email string) error
	VerifiedEmail(email string) error
	SendEmailVerification(email string)
}
