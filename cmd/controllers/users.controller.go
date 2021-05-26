package controllers

import (
	"awesomeProject/cmd/models"
	"awesomeProject/cmd/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if user.Name == "" || user.Age <= 0 || user.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "You need to provide all required fields"})
		return
	}
	c.JSON(http.StatusOK, services.CreateUser(user))
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAllUsers())
}
