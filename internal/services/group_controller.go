package services

import (
	"github.com/labstack/echo"
	"psp/internal/database"
	"psp/internal/database/interfaces"
	"psp/internal/models"
	"psp/internal/services/usecase/controller"
)

type GroupController struct {
	Interactor controller.GroupInteractor
}

func NewGroupController(sqlHandler interfaces.SqlHandler) *GroupController {
	return &GroupController{
		Interactor: controller.GroupInteractor{
			GroupRepository: &database.GroupRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *GroupController) Create(c echo.Context) {
	u := models.Groups{}
	c.Bind(&u)

	print("\ndata= ", u.UserID, " ", u.GroupID)

	controller.Interactor.Add(u)

	c.JSON(201, "")
	return
}

func (controller *GroupController) CreateByData(userId int, groupId int, roleId int) {
	u := models.Groups{}
	u.RoleID = roleId
	u.GroupID = groupId
	u.UserID = userId
	controller.Interactor.Add(u)
	return
}

func (controller *GroupController) GetGroup() []models.Groups {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *GroupController) GetGroupsForUser(userId int) []models.Groups {
	res := controller.Interactor.GetInfoForUser(userId)
	return res
}

//	func (controller *GroupController) GetGroupByName(name string) (models.Groups, bool) {
//		group, found := controller.Interactor.GetInfoByName(name)
//		return group, found
//	}
func (controller *GroupController) CheckIfExists(id int, groupId int) bool {
	found := controller.Interactor.Check(id, groupId)
	return found.UserID != 0
}

func (controller *GroupController) GetUserRole(id int, groupId int) int {
	found := controller.Interactor.Check(id, groupId)
	return found.RoleID
}

func (controller *GroupController) UpdateUserRole(userId int, groupId int, roleId int) {
	controller.Interactor.UpdateRole(userId, groupId, roleId)
}

func (controller *GroupController) Delete(id string) {
	controller.Interactor.Delete(id)
}
