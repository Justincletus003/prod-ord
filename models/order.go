package models

import "time"

type Order struct{
	Id uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	ProductRefer int `json:"product_id"`
	Product Product `gorm:"foreignKey:ProductRefer"`
	UserRefer int `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User User `gorm:"foreignKey:UserRefer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}