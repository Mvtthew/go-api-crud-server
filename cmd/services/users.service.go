package services

import (
	"awesomeProject/cmd/database"
	"awesomeProject/cmd/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(user models.User) models.User {
	database.DB.Create(&user)
	return user
}

func RemoveUserById(id uint) gin.H {
	var user models.User
	result := database.DB.Find(&user, id)
	if result.RowsAffected == 0 {
		return gin.H{"message": "User does not exists"}
	}
	database.DB.Delete(&user)
	return gin.H{"message": "User deleted"}
}

func UpdateUserById(updatedUser models.User, id uint) models.User {
	var user models.User
	result := database.DB.Find(&user, id)
	if result.RowsAffected == 0 {
		return user
	}
	result.Updates(updatedUser)
	return user
}

func GetAllUsers() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}
