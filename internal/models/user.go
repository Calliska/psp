package models

// User

type User struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Email      string `json:"email"`
	FirstName  string `json:"firstname"`
	SecondName string `json:"secondname"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
