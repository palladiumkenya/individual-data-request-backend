package services

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mailgun/mailgun-go/v4"
	"html/template"
	"log"
	"os"
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

func SendRequesterEmail(subject string, body map[string]interface{}, recipient string, email_template string, c *gin.Context) (string, error) {
	sender := "no-reply@kenyahmis.org"
	// Load environment variables
	mailgunKey := os.Getenv("MAILGUN_KEY")
	mailgunDomain := os.Getenv("MAILGUN_DOMAIN")

	if mailgunKey == "" || mailgunDomain == "" {
		return "", fmt.Errorf("mailgun key or domain is not set")
	}

	// Initialize Mailgun client
	mg := mailgun.NewMailgun(mailgunDomain, mailgunKey)

	// Parse the HTML template
	t, err := template.ParseFiles(email_template)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		return "", err
	}

	// Execute the template with the provided data
	var bod bytes.Buffer
	if err := t.Execute(&bod, body); err != nil {
		log.Printf("Error executing template: %v", err)
		return "", err
	}

	// Create a new Mailgun message with HTML content
	m := mg.NewMessage(sender, subject, "", recipient)
	m.SetHtml(bod.String()) // Set HTML body content

	// Send the email
	_, id, err := mg.Send(c, m)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return "", err
	}

	return id, err
}
