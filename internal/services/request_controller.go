package services

import (
	"github.com/labstack/echo"
	"psp/internal/database"
	"psp/internal/database/interfaces"
	"psp/internal/models"
	"psp/internal/services/usecase/controller"
)

type RequestController struct {
	Interactor controller.RequestInteractor
}

func NewRequestController(sqlHandler interfaces.SqlHandler) *RequestController {
	return &RequestController{
		Interactor: controller.RequestInteractor{
			RequestRepository: &database.RequestRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RequestController) Create(c echo.Context) {
	u := models.Request{}
	c.Bind(&u)

	controller.Interactor.Add(u)
	return
}

func (controller *RequestController) CreateByData(userId int, groupId int, status int) {
	u := models.Request{}
	u.Status = status
	u.GroupId = groupId
	u.UserId = userId
	controller.Interactor.Add(u)
	return
}

func (controller *RequestController) UpdateRequestStatus(requestId int, status int) {

	controller.Interactor.SetStatus(requestId, status)
	return
}

//func (controller *RequestController) GetGroup() []models.Groups {
//	res := controller.Interactor.GetInfo()
//	return res
//}

//	func (controller *GroupController) GetGroupByName(name string) (models.Groups, bool) {
//		group, found := controller.Interactor.GetInfoByName(name)
//		return group, found
//	}
func (controller *RequestController) CheckIfExists(UserId int, GroupId int) bool {
	found := controller.Interactor.Check(UserId, GroupId)
	return found.UserId != 0
}

func (controller *RequestController) GetRequestById(requestId int) models.Request {
	found := controller.Interactor.GetInfoById(requestId)
	return found
}

func (controller *RequestController) Delete(id string) {
	controller.Interactor.Delete(id)
}
