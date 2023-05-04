package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"psp/internal/database"
	"psp/internal/database/interfaces"
	"psp/internal/models"
	"psp/internal/services/usecase/controller"
	"time"
)

var jwtSecretKey = []byte("super-secret-what-is-it-i-do-not-know-1@@13")

type UserController struct {
	Interactor controller.UserInteractor
}

func NewUserController(sqlHandler interfaces.SqlHandler) *UserController {
	return &UserController{
		Interactor: controller.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) {
	u := models.User{}
	c.Bind(&u)

	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = string(hashedPassword)

	controller.Interactor.Add(u)

	createdUsers := controller.Interactor.GetInfoByEmail(u.Email)

	c.JSON(201, createdUsers)
	return
}

func (controller *UserController) GetUser() []models.User {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) GetUserByEmail(email string) models.User {
	res := controller.Interactor.GetInfoByEmail(email)
	return res
}

func (controller *UserController) GetUserByToken(token string) (models.User, bool) {
	user := models.User{}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(jwtToken *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return user, false
	}

	user = controller.GetUserByEmail(claims["sub"].(string))

	if len(user.Email) == 0 {
		return user, false
	}

	return user, true

}

func (controller *UserController) LoginUserByCredits(email string, password string) (models.LoginResponse, bool) {
	user := controller.GetUserByEmail(email)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return models.LoginResponse{}, false
	}

	payload := jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return models.LoginResponse{}, false
	}
	return models.LoginResponse{AccessToken: t}, true
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
