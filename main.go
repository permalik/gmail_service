package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

var (
	msg        = []byte("email message 003...")
	recipients = []string{"tymalik@pm.me"}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	from := os.Getenv("FROM")

	hostname := os.Getenv("HOSTNAME")
	auth := smtp.PlainAuth("", username, password, hostname)

	err = smtp.SendMail(hostname+":587", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}
