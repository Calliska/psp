package models

// Group

type GroupNames struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
