package services

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
	"fmt"
)

func SendSimpleMessage(sender, subject, body, recipient string, c *gin.Context) (string, error) {
	mailgun_key := os.Getenv("MAILGUN_KEY")
	mailgun_domain := os.Getenv("MAILGUN_DOMAIN")
	fmt.Printf("Mailgun Domain: %s", mailgun_domain)

	mg := mailgun.NewMailgun(mailgun_domain, mailgun_key)

	m := mg.NewMessage(sender, subject, body, recipient)

	_, id, err := mg.Send(c, m)
	return id, err
}

