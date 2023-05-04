package database

import (
	"psp/internal/database/interfaces"
	"psp/internal/models"
)

type GroupRepository struct {
	interfaces.SqlHandler
}

func (db *GroupRepository) Store(g models.Groups) {
	db.Create(&g)
}

func (db *GroupRepository) Select() []models.Groups {
	var group []models.Groups
	db.FindAll(&group)
	return group
}

func (db *GroupRepository) SelectForUser(userId int) []models.Groups {
	var group []models.Groups
	db.Preload("Groups").Where("user_id = ?", userId).Find(&group)
	return group
}

func (db *GroupRepository) SelectByName(name string) (models.Groups, bool) {
	var group models.Groups
	var groupName models.GroupNames
	db.Preload("GroupNames").Where("name = ?", name).Find(&groupName).Preload("Groups").Where("group_id = ?", groupName.Id).Find(&group)
	if len(groupName.Name) > 0 {
		return group, true
	}
	return group, false
}

func (db *GroupRepository) SelectUserById(id int, groupId int) models.Groups {
	var group models.Groups
	db.Preload("Groups").Where("user_id = ?", id).Where("group_id = ?", groupId).Find(&group)
	return group
}

func (db *GroupRepository) UpdateRoleById(userId int, groupId int, roleId int) {
	db.Preload("Group").Model(&models.Groups{}).Where("user_id = ?", userId).Where("group_id = ?", groupId).Update("role_id", roleId)
}

func (db *GroupRepository) Delete(id string) {
	var group []models.Groups
	db.DeleteById(&group, id)
}
