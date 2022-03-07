package utils

import (
	"errors"
	"fiberAPI/db"
	"fiberAPI/models"
)

func CheckUser(id int, user *models.User) error {

	db.Database.Database.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("there is no user with entered id")
	}
	return nil
}
