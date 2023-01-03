package repository

import (
	"github.com/adibSetiawann/cronjob/config"
	"github.com/adibSetiawann/cronjob/entity"
	"github.com/adibSetiawann/cronjob/model"
)

type WalletRepositoryImplement struct {
	
}

func (*WalletRepositoryImplement) Create(wallet *entity.Wallet) error{
	db := config.DB.Debug().Create(&wallet)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (*WalletRepositoryImplement) FindById(id string) ([]model.WalletResponse, error) {

	var wallets []model.WalletResponse

	err := config.DB.Debug().Preload("User").First(&wallets, "id=?", id)
	if err.Error != nil {
		return nil, err.Error
	}

	return wallets, nil
}

func(*WalletRepositoryImplement) FindAll() ([]model.WalletResponse, error) {

	var wallets []model.WalletResponse

	db := config.DB.Debug().Preload("Gender").Find(&wallets)
	if db.Error != nil {
		return nil, db.Error
	}

	return wallets, nil
}


func NewWalletRepository()  WalletRepository{
	return &WalletRepositoryImplement{}
}