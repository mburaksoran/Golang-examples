package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mburaksoran/fiberv2restapi/db"
	"github.com/mburaksoran/fiberv2restapi/user"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	database, err := db.Connect()

	if err != nil {
		log.Fatalf("cannot connect to database")
	}
	fmt.Println(database.Config)
	tracer.Start()

	defer tracer.Stop()
	repo := user.NewRepository(database)
	err = repo.Migration()
	if err != nil {
		log.Fatal(err)
	}
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	app := fiber.New()
	app.Use(fibertrace.Middleware(), logger.New())
	app.Get("/users/:id", handler.Get)
	app.Post("/users", handler.Create)
	app.Delete("/users/:id", handler.Delete)
	app.Put("/users/:id", handler.Update)
	app.Listen(":8080")
}
