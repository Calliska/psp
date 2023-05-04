package repo

import "psp/internal/models"

type GroupRepository interface {
	Store(models.Groups)
	Select() []models.Groups
	SelectForUser(userId int) []models.Groups
	SelectByName(name string) (models.Groups, bool)
	SelectUserById(id int, groupId int) models.Groups
	UpdateRoleById(userId int, groupId int, roleId int)
	Delete(id string)
}
