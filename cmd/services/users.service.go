package services

import (
	"awesomeProject/cmd/database"
	"awesomeProject/cmd/models"
)

func CreateUser(user models.User)  {
	database.DB.Create(user)
}

func GetAllUsers() []models.User  {
	var users []models.User
	database.DB.Find(&users)
	return users
}
