package entity

type Wallet struct {
	ID         int     `gorm:"primaryKey" json:"id"`
	Amount     float64 `json:"amount" validate:"required"`
	CurrencyId int     `json:"currency_id"`
	UserId     int     `json:"user_id"`
}
