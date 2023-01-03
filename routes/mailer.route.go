package routes

import (
	"github.com/adibSetiawann/cronjob/controller"
	"github.com/adibSetiawann/cronjob/middleware"
	repository "github.com/adibSetiawann/cronjob/repository/mailer"
	service "github.com/adibSetiawann/cronjob/service/mailer"
	"github.com/gofiber/fiber/v2"
)

func MailerRoute(app *fiber.App) {
	mailerRepo := repository.NewMailerRepository()
	mailerService := service.NewMailerService(&mailerRepo)
	mailerController := controller.NewMailerController(&mailerService)

	app.Post("/send-otp", middleware.AuthForRegistered,mailerController.SendEmail)
	app.Post("/verify-otp", middleware.AuthForRegistered, mailerController.VerifiyEmail)
}
