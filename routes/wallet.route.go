package routes

import (
	"github.com/adibSetiawann/cronjob/controller"
	"github.com/adibSetiawann/cronjob/middleware"
	repository "github.com/adibSetiawann/cronjob/repository/wallet"
	service "github.com/adibSetiawann/cronjob/service/wallet"
	"github.com/gofiber/fiber/v2"
)

func WalletRoute(app *fiber.App) {
	walletRepo := repository.NewWalletRepository()
	walletService := service.NewWalletService(&walletRepo)
	walletController := controller.NewWalletController(&walletService)

	app.Post("/wallets/create", walletController.Create)
	app.Get("/wallets", middleware.AuthAsAdmin, walletController.GetAll)
	app.Get("/wallets/:id", middleware.AuthForRegistered, walletController.GetById)
}
