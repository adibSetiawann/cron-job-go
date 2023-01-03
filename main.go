package main

import (
	"github.com/adibSetiawann/cronjob/routes"
	"github.com/adibSetiawann/cronjob/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config.ConnectDatabase()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "success create API",
		})
	})

	routes.UserRoute(app)
	routes.WalletRoute(app)

	app.Listen(":8080")
}
