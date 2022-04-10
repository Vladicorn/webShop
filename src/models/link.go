package models

type Link struct {
	Id       uint      `json:"id"`
	Code     string    `json:"code"`
	UserId   uint      `json:"user_id"`
	User     User      `gorm:"foreignKey:UserId"`
	Products []Product `gorm:"many2many:link_products"`
}
