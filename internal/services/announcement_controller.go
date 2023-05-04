package services

import (
	"github.com/labstack/echo"
	"psp/internal/database"
	"psp/internal/database/interfaces"
	"psp/internal/models"
	"psp/internal/services/usecase/controller"
)

type AnnouncementController struct {
	Interactor controller.AnnouncementInteractor
}

func NewAnnouncementController(sqlHandler interfaces.SqlHandler) *AnnouncementController {
	return &AnnouncementController{
		Interactor: controller.AnnouncementInteractor{
			AnnouncementRepository: &database.AnnouncementRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AnnouncementController) Create(c echo.Context, announcement models.Announcement) {
	//u := models.Announcement{}
	c.Bind(&announcement)

	//print("id=", u.Creator, "\tdate=", u.Date.String(), "\n")

	controller.Interactor.Add(announcement)

	//c.JSON(201, "")
	return
}

func (controller *AnnouncementController) GetAnnouncements(id int) []models.Announcement {
	res := controller.Interactor.GetInfo(id)
	return res
}

func (controller *AnnouncementController) Delete(id string) {
	controller.Interactor.Delete(id)
}
