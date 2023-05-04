package controller

import (
	"psp/internal/models"
	"psp/internal/services/usecase/repo"
)

type RequestInteractor struct {
	RequestRepository repo.RequestRepository
}

func (interactor *RequestInteractor) Add(u models.Request) {
	interactor.RequestRepository.Store(u)
}

func (interactor *RequestInteractor) SetStatus(requestId int, status int) {
	interactor.RequestRepository.UpdateStatus(requestId, status)
}

func (interactor *RequestInteractor) GetInfo() []models.Request {
	return interactor.RequestRepository.Select()
}

func (interactor *RequestInteractor) GetInfoById(requestId int) models.Request {
	return interactor.RequestRepository.SelectRequestById(requestId)
}

func (interactor *RequestInteractor) Check(UserId int, GroupId int) models.Request {
	return interactor.RequestRepository.SelectRequestByData(UserId, GroupId)
}

func (interactor *RequestInteractor) Delete(id string) {
	interactor.RequestRepository.Delete(id)
}
