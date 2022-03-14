package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mburaksoran/fiberv2restapi/db"
	"github.com/mburaksoran/fiberv2restapi/user"
)

func main() {
	database, err := db.Connect()

	if err != nil {
		log.Fatalf("cannot connect to database")
	}
	fmt.Println(database.Config)

	repo := user.NewRepository(database)
	err = repo.Migration()
	if err != nil {
		log.Fatal(err)
	}
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	app.Post("/users", handler.Create)
	app.Delete("/users/:id", handler.Delete)
	app.Put("/users/:id", handler.Update)
	app.Listen(":8080")
}
