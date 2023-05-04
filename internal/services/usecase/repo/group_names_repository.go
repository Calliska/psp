package repo

import "psp/internal/models"

type GroupNamesRepository interface {
	Store(models.GroupNames)
	Select() []models.GroupNames
	SelectByName(name string) models.GroupNames
	SelectById(id int) models.GroupNames
	Delete(id string)
}
