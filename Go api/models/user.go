package models

import "time"

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	CreatedAt time.Time
}
