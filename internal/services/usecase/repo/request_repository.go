package repo

import "psp/internal/models"

type RequestRepository interface {
	Store(request models.Request)
	UpdateStatus(requestId int, status int)
	Select() []models.Request
	SelectRequestByData(UserId int, GroupId int) models.Request
	SelectRequestById(requestId int) models.Request
	Delete(id string)
}
