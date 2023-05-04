package transport

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"psp/internal/database"
	"psp/internal/models"
	controllers "psp/internal/services"
	"strconv"
	"time"
)

// TODO:
// 		[+] Create group method
// 		[+] Send request to join group method
// 		[+] Create announcement method
// 		[+] Show all announcements for user method
//		[+] Get groups method
// 		[]
//		[+] Change user role method
//		[+] Accept user to the group method

func Init() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}))

	userController := controllers.NewUserController(database.NewSqlHandler())
	groupController := controllers.NewGroupController(database.NewSqlHandler())
	requestController := controllers.NewRequestController(database.NewSqlHandler())
	announcementController := controllers.NewAnnouncementController(database.NewSqlHandler())
	groupNamesController := controllers.NewGroupNamesController(database.NewSqlHandler())

	e.POST("/register", func(c echo.Context) error {
		u := models.User{}
		c.Bind(&u)
		user := userController.GetUserByEmail(u.Email)
		if len(user.Email) == 0 {
			userController.Create(c)
			return c.String(http.StatusOK, "created")
		} else {
			return c.String(http.StatusOK, "already exists")
		}
	})

	e.POST("/login", func(c echo.Context) error {
		u := models.User{}
		c.Bind(&u)
		token, ok := userController.LoginUserByCredits(u.Email, u.Password)
		if ok {
			return c.JSON(http.StatusOK, token)
		}
		return c.String(http.StatusUnauthorized, "wrong token")
	})

	e.GET("/profile/:token", func(c echo.Context) error {
		user, ok := userController.GetUserByToken(c.Param("token"))
		if ok {
			return c.JSON(http.StatusOK, user)
		}
		return c.String(http.StatusUnauthorized, "wrong token")
	})

	e.POST("/create/:token/:group_name", func(c echo.Context) error {
		token := c.Param("token")
		groupName := c.Param("group_name")
		user, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		group, ok := groupNamesController.CreateByName(groupName)
		c.Bind(&group)
		if !ok {
			return c.String(http.StatusOK, "group already exists")
		}

		groupController.CreateByData(user.Id, group.Id, 3)

		return c.String(http.StatusOK, "group created")
	})

	e.POST("/join/request/:token/:group_id", func(c echo.Context) error {
		token := c.Param("token")
		groupId, _ := strconv.Atoi(c.Param("group_id"))

		user, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		group := groupNamesController.GetGroupById(groupId)

		if len(group.Name) <= 0 {
			return c.String(http.StatusOK, "unknown group")
		}

		exists := groupController.CheckIfExists(user.Id, groupId)

		if exists {
			return c.String(http.StatusOK, "already in group")
		}

		exists = requestController.CheckIfExists(user.Id, groupId)
		if exists {
			return c.String(http.StatusOK, "request already exists")
		}

		requestController.CreateByData(user.Id, groupId, 0)

		return c.String(http.StatusOK, "request has been sent")
	})

	e.POST("/accept/request/:token/:request_id", func(c echo.Context) error {
		token := c.Param("token")
		requestId, _ := strconv.Atoi(c.Param("request_id"))

		user, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		request := requestController.GetRequestById(requestId)
		if request.Id == 0 {
			return c.String(http.StatusOK, "no such request")
		}

		groupRole := groupController.GetUserRole(user.Id, request.GroupId)

		if groupRole <= 1 {
			return c.String(http.StatusUnauthorized, "not enough rights")
		}

		if request.Status == 1 {
			return c.String(http.StatusOK, "already accepted to the group")
		}

		if request.Status == -1 {
			return c.String(http.StatusOK, "request was declined")
		}

		requestController.UpdateRequestStatus(requestId, 1)
		groupController.CreateByData(request.UserId, request.GroupId, 1)

		return c.String(http.StatusOK, "User accepted to the group")
	})

	e.POST("/decline/request/:token/:request_id", func(c echo.Context) error {
		token := c.Param("token")
		requestId, _ := strconv.Atoi(c.Param("request_id"))

		user, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		request := requestController.GetRequestById(requestId)
		if request.Id == 0 {
			return c.String(http.StatusOK, "no such request")
		}

		groupRole := groupController.GetUserRole(user.Id, request.GroupId)

		if groupRole <= 1 {
			return c.String(http.StatusUnauthorized, "not enough rights")
		}

		if request.Status == 1 {
			return c.String(http.StatusOK, "already accepted to the group")
		}

		if request.Status == -1 {
			return c.String(http.StatusOK, "already declined")
		}

		requestController.UpdateRequestStatus(requestId, -1)

		return c.String(http.StatusOK, "User accepted to the group")
	})

	e.POST("/user/changeRole/:token/:group_id/:role_id/:user_id", func(c echo.Context) error {
		token := c.Param("token")
		groupId, _ := strconv.Atoi(c.Param("group_id"))
		roleId, _ := strconv.Atoi(c.Param("role_id"))
		userId, _ := strconv.Atoi(c.Param("user_id"))

		user, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		groupRole := groupController.GetUserRole(user.Id, groupId)

		if groupRole <= 1 {
			return c.String(http.StatusUnauthorized, "not enough rights")
		}

		ok = groupController.CheckIfExists(userId, groupId)

		if !ok {
			return c.String(http.StatusOK, "user not in group")
		}

		groupController.UpdateUserRole(userId, groupId, roleId)

		return c.String(http.StatusOK, "role changed")
	})

	e.POST("/announcement/create/:token", func(c echo.Context) error {
		token := c.Param("token")

		user, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		a := models.Announcement{}
		c.Bind(&a)
		a.Creator = user.Id
		a.Date = time.Now()

		groupRole := groupController.GetUserRole(user.Id, a.GroupId)
		print("GroupId = ")

		if groupRole <= 1 {
			return c.String(http.StatusUnauthorized, "not enough rights")
		}

		announcementController.Create(c, a)

		return c.String(http.StatusOK, "created")
	})

	e.GET("/announcement/get/:token/:group_id", func(c echo.Context) error {
		token := c.Param("token")
		groupId, _ := strconv.Atoi(c.Param("group_id"))

		user, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		exists := groupController.CheckIfExists(user.Id, groupId)

		if !exists {
			return c.String(http.StatusUnauthorized, "do not belongs to group")
		}
		print("group id = ", groupId, "\n")
		result := announcementController.GetAnnouncements(groupId)
		c.Bind(&result)

		return c.JSON(http.StatusOK, result)
	})

	e.GET("/groups/get/all/:token", func(c echo.Context) error {
		token := c.Param("token")

		_, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		groups := groupNamesController.GetGroup()

		return c.JSON(http.StatusOK, groups)
	})

	e.GET("/groups/get/forUser/:token", func(c echo.Context) error {
		token := c.Param("token")

		user, ok := userController.GetUserByToken(token)
		if !ok {
			return c.String(http.StatusUnauthorized, "wrong token")
		}

		groups := groupController.GetGroupsForUser(user.Id)
		var groupNames = make([]models.GroupNames, len(groups))

		for key, value := range groups {
			groupNames[key] = groupNamesController.GetGroupById(value.GroupID)
		}

		return c.JSON(http.StatusOK, groupNames)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
