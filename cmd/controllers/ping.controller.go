package controllers

import "github.com/gin-gonic/gin"

func HandlePingPong(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "PONG",
	})
}