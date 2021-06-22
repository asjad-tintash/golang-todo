package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(email string) error {
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{
		email,
	}

	smtpHost := "smtp.gmail.com"

	auth := smtp.PlainAuth("", from, password, smtpHost)
	message := "Signup at http://localhost:5000/register with your email"
	msg := "From: " + from + "\n" +  "Subject: ToDo App\n\n" + message

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, []byte(msg))
	fmt.Println(err)
	if err != nil {
		return err
	}

	fmt.Println("Email sent successfully")

	return nil
}
