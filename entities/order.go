package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TotalPrice  uint       `gorm:"not null" json:"total_price" form:"total_price"`
	StatusOrder uint       `gorm:"not null" json:"status_order" form:"status_order"`
	Address     Address    `gorm:"foreignKey:ID;references:ID" json:"address" form:"address"`
	CreditCard  CreditCard `gorm:"foreignKey:ID;references:ID" json:"credit_card" form:"credit_card"`
	Cart        []Cart     `gorm:"foreignKey:OrderID;references:ID"`
}
