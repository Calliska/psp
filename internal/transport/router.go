package transport

import (
	"github.com/labstack/echo"
	"net/http"
	"psp/internal/database"
	controllers "psp/internal/services"
)

func Init() {
	e := echo.New()

	userController := controllers.NewUserController(database.NewSqlHandler())

	e.GET("/users", func(c echo.Context) error {
		users := userController.GetUser()
		c.Bind(&users)
		return c.JSON(http.StatusOK, users)
	})

	e.POST("/users", func(c echo.Context) error {
		userController.Create(c)
		return c.String(http.StatusOK, "created")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
