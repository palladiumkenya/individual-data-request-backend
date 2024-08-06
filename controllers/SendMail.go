package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palladiumkenya/individual-data-request-backend/models"
	"github.com/palladiumkenya/individual-data-request-backend/services"
)

func SendMail(c *gin.Context) {
	/*
		This is controller is used for sending mail
	*/
	var mail models.Mail
	if err := c.BindJSON(&mail); err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	id, err := services.SendSimpleMessage(mail.Sender, mail.Subject, mail.Body, mail.Recipient, c)

	c.JSON(http.StatusOK, id)
	fmt.Printf("Error: %s", err)
}
