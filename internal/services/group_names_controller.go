package services

import (
	"github.com/labstack/echo"
	"psp/internal/database"
	"psp/internal/database/interfaces"
	"psp/internal/models"
	"psp/internal/services/usecase/controller"
)

type GroupNamesController struct {
	Interactor controller.GroupNamesInteractor
}

func NewGroupNamesController(sqlHandler interfaces.SqlHandler) *GroupNamesController {
	return &GroupNamesController{
		Interactor: controller.GroupNamesInteractor{
			GroupNamesRepository: &database.GroupNamesRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *GroupNamesController) Create(c echo.Context) models.GroupNames {
	g := models.GroupNames{}
	c.Bind(&g)

	print("adding name: ", g.Name)
	controller.Interactor.Add(g)

	c.JSON(201, "")
	return g
}

func (controller *GroupNamesController) CreateByName(name string) (models.GroupNames, bool) {
	g := models.GroupNames{Name: name}

	findGroup := controller.GetGroupByName(name)
	if len(findGroup.Name) > 0 {
		return g, false
	}

	controller.Interactor.Add(g)

	findGroup = controller.GetGroupByName(name)

	print("Cretaed group id = ", findGroup.Id, "")

	return findGroup, true
}

func (controller *GroupNamesController) GetGroup() []models.GroupNames {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *GroupNamesController) GetGroupByName(name string) models.GroupNames {
	res := controller.Interactor.GetInfoByName(name)
	return res
}

func (controller *GroupNamesController) GetGroupById(id int) models.GroupNames {
	res := controller.Interactor.GetInfoById(id)
	return res
}

func (controller *GroupNamesController) Delete(id string) {
	controller.Interactor.Delete(id)
}
