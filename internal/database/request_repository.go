package database

import (
	"psp/internal/database/interfaces"
	"psp/internal/models"
)

type RequestRepository struct {
	interfaces.SqlHandler
}

func (db *RequestRepository) Store(g models.Request) {
	db.Create(&g)
}

func (db *RequestRepository) UpdateStatus(requestId int, status int) {
	print("Updating id=", requestId, " with status=", status, "\n")
	db.Preload("Request").Model(&models.Request{}).Where("id = ?", requestId).Update("status", status)

}

func (db *RequestRepository) Select() []models.Request {
	var group []models.Request
	db.FindAll(&group)
	return group
}

//	func (db *RequestRepository) SelectByName(name string) (models.Groups, bool) {
//		var group models.Groups
//		var groupName models.GroupNames
//		db.Preload("GroupNames").Where("name = ?", name).Find(&groupName).Preload("Groups").Where("group_id = ?", groupName.Id).Find(&group)
//		if len(groupName.Name) > 0 {
//			return group, true
//		}
//		return group, false
//	}

func (db *RequestRepository) SelectRequestById(requestId int) models.Request {
	var group models.Request
	db.Preload("Request").Where("id = ?", requestId).Find(&group)
	return group
}

func (db *RequestRepository) SelectRequestByData(UserId int, GroupId int) models.Request {
	var group models.Request
	db.Preload("Request").Where("user_id = ?", UserId).Where("group_id = ?", GroupId).Find(&group)
	return group
}

func (db *RequestRepository) Delete(id string) {
	var group []models.Groups
	db.DeleteById(&group, id)
}
