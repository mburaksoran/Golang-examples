package main

import (
	"fiberAPI/db"
	"fiberAPI/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func welcome(f *fiber.Ctx) error {
	return f.SendString("welcome to fiber api")
}

func main() {
	db.DbMigration()
	app := fiber.New()
	app.Get("/", welcome)
	app.Post("/api/createuser", handlers.CreateUser)
	app.Get("/api/getuser", handlers.GetUsers)
	app.Get("/api/getuser/:id", handlers.GetUserById)
	app.Delete("/api/deleteuser/:id", handlers.DeleteUser)
	app.Put("/api/updateuser/:id", handlers.UpdateUser)
	log.Fatal(app.Listen(":8080"))

}
