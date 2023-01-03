package middleware

import (
	"github.com/adibSetiawann/cronjob/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthForRegistered(ctx *fiber.Ctx) error {
	token := ctx.Cookies("jwt")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	_, err := utils.VerifyToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	return ctx.Next()
}

func AuthAsAdmin(ctx *fiber.Ctx) error {
	token := ctx.Cookies("jwt")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access for this token",
		})
	}

	ctx.Locals("role", claims["role"])

	return ctx.Next()
}

func AuthAsUser(ctx *fiber.Ctx) error {
	token := ctx.Cookies("jwt")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	role := claims["role"].(string)
	if role != "User" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access for this token",
		})
	}

	ctx.Locals("role", claims["role"])

	return ctx.Next()
}
