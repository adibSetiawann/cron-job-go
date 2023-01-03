package service

import "github.com/adibSetiawann/cronjob/model"

type WalletService interface {
	Create(customerRequest model.CreateWallet) (model.WalletResponse, error)
	GetAllData() ([]model.WalletResponse, error)
	GetById(id string) ([]model.WalletResponse, error)
	Validation(customerRequest interface{}) error
}