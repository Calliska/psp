package controller

import (
	"psp/internal/models"
	"psp/internal/services/usecase/repo"
)

type GroupInteractor struct {
	GroupRepository repo.GroupRepository
}

func (interactor *GroupInteractor) Add(u models.Groups) {
	interactor.GroupRepository.Store(u)
}

func (interactor *GroupInteractor) GetInfo() []models.Groups {
	return interactor.GroupRepository.Select()
}

func (interactor *GroupInteractor) GetInfoForUser(userId int) []models.Groups {
	return interactor.GroupRepository.SelectForUser(userId)
}

func (interactor *GroupInteractor) Check(id int, groupId int) models.Groups {
	return interactor.GroupRepository.SelectUserById(id, groupId)
}

func (interactor *GroupInteractor) UpdateRole(userId int, groupId int, roleId int) {
	interactor.GroupRepository.UpdateRoleById(userId, groupId, roleId)
}

func (interactor *GroupInteractor) Delete(id string) {
	interactor.GroupRepository.Delete(id)
}
