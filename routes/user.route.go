package routes

import (
	"github.com/adibSetiawann/cronjob/controller"
	"github.com/adibSetiawann/cronjob/middleware"
	repository "github.com/adibSetiawann/cronjob/repository/user"
	service "github.com/adibSetiawann/cronjob/service/user"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(&userRepo)
	userController := controller.NewUserController(&userService)

	app.Post("/signup", userController.Signup)
	app.Post("/signin", userController.Signin)
	app.Post("/logout", middleware.AuthForRegistered, userController.Logout)
	app.Get("/users", middleware.AuthAsAdmin, userController.GetAll)
	app.Get("/users/:id", middleware.AuthForRegistered, userController.GetById)
	app.Post("/users/update", middleware.AuthAsUser, userController.UpdateEmail)
}
