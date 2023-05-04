package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"psp/internal/database/interfaces"
	"psp/internal/models"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() interfaces.SqlHandler {
	dbServer := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbServer), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}
func (handler *SqlHandler) Create(obj interface{}) {
	handler.db.Create(obj)
}
func (handler *SqlHandler) FindAll(obj interface{}) {
	handler.db.Find(obj)
}
func (handler *SqlHandler) DeleteById(obj interface{}, id string) {
	handler.db.Delete(obj, id)
}
func (handler *SqlHandler) SelectById(obj interface{}, id string) {
	handler.db.Select(obj, id)
}
func (handler *SqlHandler) Where(object interface{}, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Where(object, args)
}
func (handler *SqlHandler) Preload(query string, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Preload(query, args)
}

func (handler *SqlHandler) Update(column string, obj interface{}) {
	handler.db.Model(&models.Request{}).Update(column, obj)
}
