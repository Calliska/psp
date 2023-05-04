package controller

import (
	"psp/internal/models"
	"psp/internal/services/usecase/repo"
)

type AnnouncementInteractor struct {
	AnnouncementRepository repo.AnnouncementRepository
}

func (interactor *AnnouncementInteractor) Add(u models.Announcement) {
	interactor.AnnouncementRepository.Store(u)
}

func (interactor *AnnouncementInteractor) GetInfo(id int) []models.Announcement {
	return interactor.AnnouncementRepository.SelectById(id)
}

func (interactor *AnnouncementInteractor) Delete(id string) {
	interactor.AnnouncementRepository.Delete(id)
}
