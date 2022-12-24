package provider

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(to string, subject string, body string) error {
	from := os.Getenv("EMAIL_USERNAME")
	password := os.Getenv("EMAIL_PASSWORD")

	toEmail := []string{to}

	host := os.Getenv("EMAIL_HOST")
	port := os.Getenv("EMAIL_PORT")
	address := host + ":" + port

	msg := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, toEmail, msg)

	if err != nil {
		log.Println(err)
	}

	return err
}