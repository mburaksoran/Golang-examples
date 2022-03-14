package db

import (
	"github.com/mburaksoran/fiberv2restapi/shared"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(shared.Config.POSTGRESURL), &gorm.Config{})
}
