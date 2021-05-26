package services

import (
	"awesomeProject/cmd/database"
	"awesomeProject/cmd/models"
)

func CreateUser(user models.User) models.User {
	database.DB.Create(&user)
	return user
}

func GetAllUsers() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}
