package main

import (
    "fmt"
	"log"
	"net/smtp"
	"os"
    "time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	from := os.Getenv("FROM")
    today := time.Now().Date()
    msg        = []byte(fmt.Sprintf("Redeployment for %s:\n- archive_ui\n- archive_resume", today))
	recipients = []string{"tymalik@pm.me"}

	hostname := os.Getenv("HOSTNAME")
	auth := smtp.PlainAuth("", username, password, hostname)

	err = smtp.SendMail(hostname+":587", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}
}
