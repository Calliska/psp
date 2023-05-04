package controller

import (
	"psp/internal/models"
	"psp/internal/services/usecase/repo"
)

type GroupNamesInteractor struct {
	GroupNamesRepository repo.GroupNamesRepository
}

func (interactor *GroupNamesInteractor) Add(u models.GroupNames) {
	interactor.GroupNamesRepository.Store(u)
}

func (interactor *GroupNamesInteractor) GetInfo() []models.GroupNames {
	return interactor.GroupNamesRepository.Select()
}

func (interactor *GroupNamesInteractor) GetInfoByName(name string) models.GroupNames {
	return interactor.GroupNamesRepository.SelectByName(name)
}

func (interactor *GroupNamesInteractor) GetInfoById(id int) models.GroupNames {
	return interactor.GroupNamesRepository.SelectById(id)
}

func (interactor *GroupNamesInteractor) Delete(id string) {
	interactor.GroupNamesRepository.Delete(id)
}
