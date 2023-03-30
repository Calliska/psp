package models

import "time"

// Announcement

type Announcement struct {
	Id      int       `json:"id" gorm:"primary_key"`
	Group   int       `json:"group"`
	Text    string    `json:"text"`
	Date    time.Time `json:"date"`
	Creator int       `json:"creator"`
}
