package helpers

import (
	"fmt"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

func SendMail(to []string, cc []string, subject, message string) error {
	CONFIG_SMTP_HOST := os.Getenv("CONFIG_SMTP_HOST")         //"smtp.gmail.com"
	CONFIG_SMTP_PORT := os.Getenv("CONFIG_SMTP_PORT")         //587
	CONFIG_SENDER_NAME := os.Getenv("CONFIG_SENDER_NAME")     //"PT. Makmur Subur Jaya <emailanda@gmail.com>"
	CONFIG_AUTH_EMAIL := os.Getenv("CONFIG_AUTH_EMAIL")       //"emailanda@gmail.com"
	CONFIG_AUTH_PASSWORD := os.Getenv("CONFIG_AUTH_PASSWORD") ///"passwordemailanda"

	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	port, _ := strconv.Atoi(CONFIG_SMTP_PORT)

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, port)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
