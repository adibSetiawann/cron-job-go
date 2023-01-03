package controller

import (
	"github.com/adibSetiawann/cronjob/model"
	"github.com/adibSetiawann/cronjob/service/mailer"
	"github.com/gofiber/fiber/v2"
)

type mailerController struct {
	mailerService service.MailerService
}

func NewMailerController(mailerService *service.MailerService) mailerController {
	return mailerController{mailerService: *mailerService}
}

func (mc *mailerController) SendEmail(c *fiber.Ctx) error {
	mailerReq := new(model.SendOtp)

	if err := c.BodyParser(mailerReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.mailerService.Validation(*mailerReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	errCreate := mc.mailerService.SendEmail(mailerReq)
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "OTP send successfully",
	})
}

func (mc *mailerController) VerifiyEmail(c *fiber.Ctx) error {

	payloads := new(model.VerifyEmail)
	// var status int

	if err := c.BodyParser(payloads); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.mailerService.Validation(*payloads)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	err := mc.mailerService.VerifyEmail(payloads)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "mailer not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": "OTP Verified Successfully",
	})
}