package model

import "time"

type Task struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Time      int       `json:"time"`
	Rep       int       `json:"rep"`
	Set_Count int       `json:"set_count"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
