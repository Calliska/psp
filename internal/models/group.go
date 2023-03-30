package models

// Group

type Group struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
