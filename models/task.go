package models

import (
	"time"
)

type Kontak struct {
	//ID         uint      `json:"id" gorm:"primary_key"`
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
