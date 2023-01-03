package main

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/adibSetiawann/cronjob/config"
	"github.com/adibSetiawann/cronjob/entity"
	"github.com/adibSetiawann/cronjob/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gopkg.in/gomail.v2"
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
	routes.MailerRoute(app)

	go Cron()
	go ExpiredEmail()

	app.Listen(":8080")
}

func ExpiredEmail() {
	config.ConnectDatabase()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	for {
		time.Sleep(time.Second * 5)
		var mailer []entity.Mailer
		fmt.Println("check expired email")
		config.DB.Find(&mailer, "status=?", "pending")
		for i, v := range mailer {
			if v.CreatedAt.Add(time.Minute * 4) == time.Now() {
				mailer[i].Status = "Expired"
				config.DB.Save(mailer[i])
				fmt.Println("user email is expired")
			}
		}
	}
}

func Cron() {
	config.ConnectDatabase()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	for {
		time.Sleep(time.Second * 10)
		var user []entity.User
		fmt.Println("Sending an email to user with pending status")
		config.DB.Find(&user, "status=?", "pending")
		for _, v := range user {
			SendEmailReminder(&v.Email)
		}
	}
}

func SendEmailReminder(email *string) {
	email_admin := config.GetEnvVariable("GMAIL")
	password := config.GetEnvVariable("PASS_GMAIL")
	var body bytes.Buffer
	t, _ := template.ParseFiles("./reminder.html")

	t.Execute(&body, struct{ Name string }{Name: "Pelanggan"})

	m := gomail.NewMessage()
	m.SetHeader("From", email_admin)
	m.SetHeader("To", *email)
	m.SetHeader("subject", "verify email!")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer("smtp.gmail.com", 587, email_admin, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}