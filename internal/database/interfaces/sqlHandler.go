package interfaces

import "gorm.io/gorm"

type SqlHandler interface {
	Create(object interface{})
	Where(object interface{}, conds ...interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	FindAll(object interface{})
	DeleteById(object interface{}, id string)
	SelectById(object interface{}, id string)
}
