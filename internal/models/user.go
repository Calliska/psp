package models

// User

type User struct {
	Id         int    `json:"id" gorm:"primary_key"`
	FirstName  string `json:"first-name"`
	SecondName string `json:"second-name"`
}
