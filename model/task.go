package model

type Task struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Time      int    `json:"time"`
	Rep       int    `json:"rep"`
	Set_Count int    `json:"set_count"`
	Comment   string `json:"comment"`
}
