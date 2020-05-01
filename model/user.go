package model

import "time"

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	NAME      string    `json:"name"`
	EMAIL     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
