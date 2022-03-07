package handlers

import (
	"fiberAPI/db"
	"fiberAPI/models"
	"fiberAPI/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(f *fiber.Ctx) error {
	var user models.User
	if err := f.BodyParser(&user); err != nil {
		return f.Status(400).JSON(err.Error())
	}

	db.Database.Database.Create(&user)
	return f.Status(200).JSON(user)
}

func GetUsers(f *fiber.Ctx) error {
	users := []models.User{}
	db.Database.Database.Find(&users)
	return f.Status(200).JSON(users)
}

func GetUserById(f *fiber.Ctx) error {
	id, err := f.ParamsInt("id")
	var user models.User

	if err != nil {
		return f.Status(400).JSON("invalid id, wrong type")
	}

	err = utils.CheckUser(id, &user)

	if err != nil {
		return f.Status(400).JSON(err.Error())
	}

	return f.Status(200).JSON(user)
}

func DeleteUser(f *fiber.Ctx) error {
	id, err := f.ParamsInt("id")
	var user models.User
	if err != nil {
		return f.Status(400).JSON("Wrong type")
	}
	err = utils.CheckUser(id, &user)
	if err != nil {
		return f.Status(400).JSON(err.Error())
	}
	if err = db.Database.Database.Delete(&user).Error; err != nil {
		return f.Status(404).JSON(err.Error())
	}

	return f.Status(200).JSON("Successfully deleted User " + strconv.Itoa(id))
}

func UpdateUser(f *fiber.Ctx) error {

	id, err := f.ParamsInt("id")
	var user models.User
	var update models.User

	if err != nil {
		return f.Status(400).JSON("Wrong type")
	}

	err = utils.CheckUser(id, &user)

	if err != nil {
		return f.Status(400).JSON(err.Error())
	}
	if err := f.BodyParser(&update); err != nil {
		return f.Status(500).JSON(err.Error())
	}
	user.Email = update.Email
	user.FirstName = update.FirstName
	user.LastName = update.LastName
	db.Database.Database.Save(user)
	return f.Status(200).JSON(user)
}
