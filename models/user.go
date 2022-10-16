package models

import "time"

type User struct {
	Id uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}