package models

// User

type Request struct {
	Id      int `json:"id" gorm:"primary_key"`
	UserId  int `json:"userId"`
	GroupId int `json:"groupId"`
	Status  int `json:"status"`
}
