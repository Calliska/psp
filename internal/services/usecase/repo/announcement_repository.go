package repo

import "psp/internal/models"

type AnnouncementRepository interface {
	Store(models.Announcement)
	Select() []models.Announcement
	SelectById(id int) []models.Announcement
	Delete(id string)
}
