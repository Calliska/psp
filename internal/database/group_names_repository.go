package database

import (
	"psp/internal/database/interfaces"
	"psp/internal/models"
)

type GroupNamesRepository struct {
	interfaces.SqlHandler
}

func (db *GroupNamesRepository) Store(g models.GroupNames) {
	db.Create(&g)
}

func (db *GroupNamesRepository) Select() []models.GroupNames {
	var group []models.GroupNames
	db.FindAll(&group)
	return group
}

func (db *GroupNamesRepository) SelectByName(name string) models.GroupNames {
	var group models.GroupNames
	db.Preload("GroupNames").Where("name = ?", name).First(&group)
	return group
}

func (db *GroupNamesRepository) SelectById(id int) models.GroupNames {
	var group models.GroupNames
	db.Preload("GroupNames").Where("id = ?", id).First(&group)
	return group
}

func (db *GroupNamesRepository) Delete(id string) {
	var group []models.GroupNames
	db.DeleteById(&group, id)
}
