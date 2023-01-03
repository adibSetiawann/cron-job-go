package repository

import (
	"github.com/adibSetiawann/cronjob/entity"
	"github.com/adibSetiawann/cronjob/model"
)

type WalletRepository interface {
	Create(customer *entity.Wallet) error
	FindAll() ([]model.WalletResponse, error)
	FindById(id string) ([]model.WalletResponse, error)
}
