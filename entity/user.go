package entity

type User struct {
	ID        int    `gorm:"primaryKey" form:"id" json:"id"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"-" gorm:"column:password"`
	Status    string `json:"status"`
}
