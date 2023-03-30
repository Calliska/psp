package models

// Role

type Role struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
