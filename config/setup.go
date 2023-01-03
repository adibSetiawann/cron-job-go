package config

import (
	"fmt"

	"github.com/adibSetiawann/cronjob/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := GetEnvVariable("DB_HOST")
	dbUser := GetEnvVariable("DB_USER")
	dbPass := GetEnvVariable("DB_PASSWORD")
	dbName := GetEnvVariable("DB_NAME")
	dbPort := GetEnvVariable("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database := db.AutoMigrate(
		&entity.User{},
		&entity.Currency{},
		&entity.Mailer{},
		&entity.Wallet{},
	)
	currencies := []entity.Currency{
        {ID: 1, Currency: "IDR"},
        {ID: 2, Currency: "USD"},
        {ID: 3, Currency: "EUR"},
        {ID: 4, Currency: "JPY"},
    }
    for index := range currencies {
        db.Create(&currencies[index])
    }
	if database != nil {
		fmt.Println("Can't running migration")
	}

	DB = db
}
