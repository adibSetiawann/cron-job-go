package controller

import (
	"time"

	"github.com/adibSetiawann/cronjob/model"
	"github.com/adibSetiawann/cronjob/service/user"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService *service.UserService) userController {
	return userController{userService: *userService}
}

func (mc *userController) Signin(c *fiber.Ctx) error {
	userReq := new(model.LoginForm)
	var status int

	if err := c.BodyParser(userReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.userService.Validation(*userReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	token, errToken := mc.userService.Signin(userReq)
	if errToken != nil {
		c.JSON(fiber.Map{
			"error":   errToken,
			"message": token,
		})
	}

	if token == "please input correct password" {
		return c.Status(404).JSON(fiber.Map{
			"message": token,
		})
	} else {
		status = 200
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(status).JSON(fiber.Map{
		"token": "success login",
	})
}

func (mc *userController) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func (mc *userController) Signup(c *fiber.Ctx) error {
	userReq := new(model.CreateUser)

	if err := c.BodyParser(userReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.userService.Validation(*userReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	user, errCreate := mc.userService.Signup(*userReq)
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "create data successfully",
		"data":    user,
	})
}


func (mc *userController) UpdateEmail(c *fiber.Ctx) error {
	userReq := new(model.LoginForm)

	if err := c.BodyParser(userReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.userService.Validation(*userReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	errCreate := mc.userService.UpdateEmail(*userReq)
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "update data seuccess, please verify your email",
	})
}

func (mc *userController) GetById(c *fiber.Ctx) error {
	userId := c.Params("id")

	users, err := mc.userService.GetById(userId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": users,
	})
}

func (mc *userController) GetAll(c *fiber.Ctx) error {
	users, _ := mc.userService.GetAllData()

	return c.Status(200).JSON(fiber.Map{
		"data": users,
	})
}
