package repository

import (
	"bytes"
	"math/rand"
	"strings"
	"text/template"
	"time"

	"github.com/adibSetiawann/cronjob/config"
	"github.com/adibSetiawann/cronjob/entity"
	"github.com/adibSetiawann/cronjob/model"
	"github.com/twilio/twilio-go"

	verify "github.com/twilio/twilio-go/rest/verify/v2"
	"gopkg.in/gomail.v2"
)

type MailerRepositoryImplement struct {
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func GenerateOTP() string {
	rndmString := strings.ToUpper(String(6))

	var str strings.Builder
	str.WriteString(rndmString)

	return str.String()
}

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: config.GetEnvVariable("SID_ACCOUNT_TWILIO"),
	Password: config.GetEnvVariable("AUTH_TWILIO"),
})

func (*MailerRepositoryImplement) SendOtp(user *model.SendOtp) (string, error) {
	email_admin := config.GetEnvVariable("GMAIL")
	password := config.GetEnvVariable("PASS_GMAIL")
	email := entity.Mailer{
		Email:  user.Email,
		Status: "pending",
		UserId: user.UserId,
	}

	dbEmail := config.DB.Debug().Create(&email)
	if dbEmail.Error != nil {
		return "", dbEmail.Error
	}
	var body bytes.Buffer
	t, _ := template.ParseFiles("./body.html")

	t.Execute(&body, struct{ Code string }{Code: GenerateOTP()})

	m := gomail.NewMessage()
	m.SetHeader("From", email_admin)
	m.SetHeader("To", user.Email)
	m.SetHeader("subject", "OTP Password")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer("smtp.gmail.com", 587, email_admin, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return "success", nil
}

func (*MailerRepositoryImplement) VerifiedEmail(email string) error {
	var mailer entity.Mailer
	var user entity.User

	err := config.DB.Debug().First(&mailer, "email=?", email)
	if err.Error != nil {
		return err.Error
	}

	if mailer.Status == "pending" {
		mailer.Status = "verified"

		errUpdate := config.DB.Debug().Save(&mailer).Error
		if errUpdate != nil {
			return errUpdate
		}
	}
	errUser := config.DB.Debug().First(&user, "email=?", email)
	if errUser.Error != nil {
		return errUser.Error
	}

	if user.Status == "pending" {
		user.Status = "verified"
		errUserUpdate := config.DB.Debug().Save(&user).Error
		if errUserUpdate != nil {
			return errUserUpdate
		}
		return nil
	}

	return nil
}

func (*MailerRepositoryImplement) ExpireLink(email string) error {
	var mailer entity.Mailer

	err := config.DB.Debug().First(&mailer, "email=?", email)
	if err.Error != nil {
		return err.Error
	}

	mailer.Status = "expired"

	errUpdate := config.DB.Debug().Save(&mailer).Error
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (*MailerRepositoryImplement) SendEmailVerification(email string) {
	email_admin := config.GetEnvVariable("GMAIL")
	password := config.GetEnvVariable("PASS_GMAIL")
	
	var body bytes.Buffer
	t, _ := template.ParseFiles("./body.html")

	t.Execute(&body, struct{ Code string }{Code: GenerateOTP()})

	m := gomail.NewMessage()
	m.SetHeader("From", email_admin)
	m.SetHeader("To", email)
	m.SetHeader("subject", "OTP Password")
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer("smtp.gmail.com", 587, email_admin, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func (*MailerRepositoryImplement) VerifyEmail(user *model.VerifyEmail) error {
	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(user.Email)
	params.SetCode(user.Pin)

	resp, err := client.VerifyV2.CreateVerificationCheck(config.GetEnvVariable("SID_SERVICE_TWILIO"), params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return nil
}

func NewMailerRepository() MailerRepository {
	return &MailerRepositoryImplement{}
}
