package db

import (
	"log"

	"github.com/jackc/pgx/v4/stdlib"
	"github.com/mburaksoran/fiberv2restapi/shared"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
	gormtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorm.io/gorm.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	sqltrace.Register("pgx", &stdlib.Driver{}, sqltrace.WithServiceName("my-service"))
	sqlDb, err := sqltrace.Open("pgx", shared.Config.POSTGRESURL)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gormtrace.Open(postgres.New(postgres.Config{Conn: sqlDb}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db, err

}
