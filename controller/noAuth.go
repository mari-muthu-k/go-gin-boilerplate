package controller

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func SignUp(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "ok",
	})
}