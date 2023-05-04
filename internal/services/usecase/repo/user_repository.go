package repo

import "psp/internal/models"

type UserRepository interface {
	Store(models.User)
	Select() []models.User
	SelectByEmail(email string) models.User
	Delete(id string)
}
