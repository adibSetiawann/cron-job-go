package controller

import (
	"github.com/adibSetiawann/cronjob/model"
	"github.com/adibSetiawann/cronjob/service/wallet"
	"github.com/gofiber/fiber/v2"
)

type walletController struct {
	walletService service.WalletService
}

func NewWalletController(walletService *service.WalletService) walletController {
	return walletController{walletService: *walletService}
}

func (mc *walletController) Create(c *fiber.Ctx) error {
	walletReq := new(model.CreateWallet)

	if err := c.BodyParser(walletReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.walletService.Validation(*walletReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	wallet, errCreate := mc.walletService.Create(*walletReq)
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "create wallet successfully",
		"data":    wallet,
	})
}

func (mc *walletController) GetById(c *fiber.Ctx) error {
	walletId := c.Params("id")

	wallets, err := mc.walletService.GetById(walletId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "wallet not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": wallets,
	})
}

func (mc *walletController) GetAll(c *fiber.Ctx) error {
	wallets, _ := mc.walletService.GetAllData()

	return c.Status(200).JSON(fiber.Map{
		"data": wallets,
	})
}
