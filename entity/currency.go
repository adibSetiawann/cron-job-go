package entity

type Currency struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Currency string `json:"currency"`
}
