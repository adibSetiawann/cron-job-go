package model

import "github.com/adibSetiawann/cronjob/entity"

type CreateWallet struct {
	Amount     float64 `json:"amount" validate:"required"`
	CurrencyId int     `json:"currency_id" validate:"required"`
	UserId     int     `json:"user_id" validate:"required"`
}

type WalletResponse struct {
	ID         int                  `gorm:"primaryKey" json:"id"`
	Amount     float64              `json:"amount" validate:"required"`
	CurrencyId int                  `json:"currency_id"`
	UserId     int                  `json:"user_id"`
	Currency   entity.Currency      `json:"currencies"`
	User       UserRelationResponse `json:"users"`
}

type WalletRelationResponse struct {
	ID         int             `gorm:"primaryKey" json:"id"`
	Amount     float64         `json:"amount" validate:"required"`
	// CurrencyId int             `json:"currency_id"`
	// Currency   entity.Currency `json:"currencies"`
}

func (WalletResponse) TableName() string {
	return "wallets"
}

func (CreateWallet) TableName() string {
	return "wallets"
}

func (WalletRelationResponse) TableName() string {
	return "wallets"
}
