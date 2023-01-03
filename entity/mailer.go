package entity

import "time"

type Mailer struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Email     string `json:"email"`
	Pin       string `json:"pin"`
	Status    string `json:"status"`
	CreatedAt time.Time
}
