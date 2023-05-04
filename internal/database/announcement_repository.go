package database

import (
	"psp/internal/database/interfaces"
	"psp/internal/models"
)

type AnnouncementRepository struct {
	interfaces.SqlHandler
}

func (db *AnnouncementRepository) Store(u models.Announcement) {
	db.Create(&u)
}

func (db *AnnouncementRepository) Select() []models.Announcement {
	var user []models.Announcement
	db.FindAll(&user)
	return user
}

func (db *AnnouncementRepository) SelectById(id int) []models.Announcement {
	var user []models.Announcement
	print("group id = ", id, "\n")
	db.Preload("Announcement").Where("group_id = ?", id).Find(&user)
	return user
}

func (db *AnnouncementRepository) Delete(id string) {
	var user []models.Announcement
	db.DeleteById(&user, id)
}
