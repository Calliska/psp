package controller

import (
	"psp/internal/models"
	"psp/internal/services/usecase/repo"
)

type UserInteractor struct {
	UserRepository repo.UserRepository
}

func (interactor *UserInteractor) Add(u models.User) {
	interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetInfo() []models.User {
	return interactor.UserRepository.Select()
}

func (interactor *UserInteractor) Delete(id string) {
	interactor.UserRepository.Delete(id)
}
