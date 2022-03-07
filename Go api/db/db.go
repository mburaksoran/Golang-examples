package db

import (
	"fiberAPI/models"
	"fiberAPI/shared"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open(postgres.Open(shared.Config.POSTGRESURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection succesful")
	Database = DbIns{Database: db}

}

type DbIns struct {
	Database *gorm.DB
}

var Database DbIns

func DbMigration() {
	Database.Database.AutoMigrate(&models.User{})
}
