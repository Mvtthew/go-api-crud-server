package controllers

import (
	"awesomeProject/cmd/models"
	"awesomeProject/cmd/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context)  {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, err)
	}
	c.JSON(200, user)
}

func GetAllUsers(c *gin.Context) {
	c.JSON(200, services.GetAllUsers())
}