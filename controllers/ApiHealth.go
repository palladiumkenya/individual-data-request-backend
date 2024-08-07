package controllers

import "github.com/gin-gonic/gin"

func GetApiHealth(c *gin.Context) {
	/*
		This is a test controller
	*/
	c.JSON(200, "Server is up and running!")
}
