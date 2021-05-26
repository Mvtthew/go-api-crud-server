package controllers

import (
	"awesomeProject/cmd/models"
	"awesomeProject/cmd/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func DeleteUser(c *gin.Context)  {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusAccepted, services.RemoveUserById(uint(id)))
}

func UpdateUser(c *gin.Context)  {
	id, idErr := strconv.ParseUint(c.Param("id"), 10, 64)
	if idErr != nil {
		c.JSON(http.StatusBadRequest, idErr.Error())
	}
	var user models.User
	jsonErr := c.ShouldBindJSON(&user)
	if jsonErr != nil {
		c.JSON(http.StatusBadRequest, jsonErr.Error())
		return
	}
	c.JSON(http.StatusAccepted, services.UpdateUserById(user, uint(id)))
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAllUsers())
}
