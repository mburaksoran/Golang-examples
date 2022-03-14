package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open("host=localhost user=postgres password=1144 port=5432 dbname=fiberapi sslmode=disable"), &gorm.Config{})
}
