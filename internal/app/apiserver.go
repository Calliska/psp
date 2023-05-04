package app

import (
	"fmt"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"psp/internal/config"
	"psp/internal/models"
	"psp/internal/transport"
)

type Server struct {
	state bool
}

func (server *Server) Start() {
	dbinit()
	transport.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(config.ServerPort))
}

func dbinit() {
	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{})

	err = db.Migrator().CreateTable(models.User{})
	if err != nil {
		fmt.Print("User already exists")
	}
	err = db.Migrator().CreateTable(models.Role{})
	if err != nil {
		fmt.Print("Role already exists")
	}
	err = db.Migrator().CreateTable(models.GroupNames{})
	if err != nil {
		fmt.Print("Group already exists")
	}
	err = db.Migrator().CreateTable(models.Groups{})
	if err != nil {
		fmt.Print("Groups already exists")
	}
	err = db.Migrator().CreateTable(models.Announcement{})
	if err != nil {
		fmt.Print("Announcement already exists")
	}
	err = db.Migrator().CreateTable(models.Request{})
	if err != nil {
		fmt.Print("Requests already exists")
	}
}
