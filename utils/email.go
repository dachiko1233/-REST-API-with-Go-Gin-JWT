package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendVerificationEmail(toEmail, token string) error {

	form := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	fmt.Println("Sending from:", form) //
	fmt.Println("Password length:", len(password))

	subject := "Verify you email"
	link := fmt.Sprintf("http://localhost:8080/api/v1/verify?token=%s", token)
	body := fmt.Sprintf("Click this likn to verify you email:\n\n%s", link)

	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", form, toEmail, subject, body)

	auth := smtp.PlainAuth("", form, password, "smtp.gmail.com")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		form,
		[]string{toEmail},
		[]byte(message),
	)

	if err != nil {
		fmt.Println("SMTP Error:", err)
	}
	return err
}
