package models

// User

type User struct {
	Id         int    `json:"id" gorm:"primary_key"`
	FirstName  string `json:"firstname"`
	SecondName string `json:"secondname"`
	Password   string `json:"password"`
}
