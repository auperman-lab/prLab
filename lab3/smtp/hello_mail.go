package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/smtp"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	password := os.Getenv("MAIl_PASSWORD")
	if password == "" {
		fmt.Println("Error: EMAIL_PASSWORD environment variable is not set")
		return
	}
	fmt.Printf("password:%s\n", password)

	const (
		from     = "macriidani@gmail.com"
		to       = "eleonorajosu17@gmail.com"
		smtpHost = "smtp.gmail.com"
	)
	message := []byte("Subject: Test mail\r\n\r\nHello Mail\r\n")
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err = smtp.SendMail(smtpHost+":587", auth, from, []string{to}, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully")
}
