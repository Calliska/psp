package repo

import "psp/internal/models"

type UserRepository interface {
	Store(models.User)
	Select() []models.User
	Delete(id string)
}
