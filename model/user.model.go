package model

type CreateUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type UpdateUserEmail struct {
	Email string `json:"email" validate:"required"`
}

type UserResponse struct {
	ID      int                      `gorm:"primaryKey" form:"id" json:"id"`
	Email   string                   `json:"email"`
	Status  string                   `json:"status"`
	Wallets []WalletRelationResponse `json:"wallets"`
	Mailers []MailerRelationResponse `json:"mailers"`
}

type UserRelationResponse struct {
	ID    int    `gorm:"primaryKey" form:"id" json:"id"`
	Email string `json:"email" validate:"required"`
}

type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (UserResponse) TableName() string {
	return "users"
}

func (UserRelationResponse) TableName() string {
	return "users"
}
