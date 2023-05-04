package models

import "time"

// Announcement

type Announcement struct {
	Id      int       `json:"id" gorm:"primary_key"`
	GroupId int       `json:"group_id"`
	Text    string    `json:"text"`
	Date    time.Time `json:"date"`
	Creator int       `json:"creator"`
}
