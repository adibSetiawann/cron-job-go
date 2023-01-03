package model

type SendOtp struct {
	Email string `json:"email"`
}

type VerifyEmail struct {
	Email string `json:"email" validate:"required"`
	Pin   string `json:"pin"  validate:"required"`
}
