package service

import (
	"errors"
	"log"

	"github.com/adibSetiawann/cronjob/entity"
	"github.com/adibSetiawann/cronjob/model"
	"github.com/adibSetiawann/cronjob/repository/wallet"
	"github.com/go-playground/validator/v10"
)

type WalletServiceImpl struct {
	walletRepo repository.WalletRepository
}


func (cs *WalletServiceImpl) Create(request model.CreateWallet) (model.WalletResponse, error) {
	var userResponse model.WalletResponse

	user := entity.Wallet{
		Amount:    request.Amount,
		UserId: request.UserId,
		CurrencyId: request.CurrencyId,
	}

	errCreate := cs.walletRepo.Create(&user)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return userResponse, errCreate
	}

	userResponse.ID = user.ID
	userResponse.Amount = user.Amount
	userResponse.UserId = user.UserId
	userResponse.CurrencyId = user.CurrencyId
	return userResponse, nil
}

func (ms *WalletServiceImpl) GetById(id string) ([]model.WalletResponse, error) {
	merchants, errFind := ms.walletRepo.FindById(id)
	if errFind != nil {
		return nil, errFind
	}
	return merchants, nil
}

func (cs *WalletServiceImpl) GetAllData() ([]model.WalletResponse, error) {
	users, errFind := cs.walletRepo.FindAll()
	if errFind != nil {
		return nil, errFind
	}
	return users, nil
}

func (ms *WalletServiceImpl) Validation(userRequest interface{}) error {
	var messageError string
	var isError bool

	// VALIDATION USER INPUT
	validate := validator.New()
	errValidate := validate.Struct(userRequest)
	if errValidate != nil {
		messageError += errValidate.Error()
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}

func NewWalletService(walletRepo *repository.WalletRepository) WalletService {
	return &WalletServiceImpl{
		walletRepo: *walletRepo,
	}
}
