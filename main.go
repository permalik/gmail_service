package main

import (
	"log"
	"net/smtp"
	"os"
)

var (
	username   = os.Getenv("USERNAME")
	password   = os.Getenv("PASSWORD")
	from       = os.Getenv("FROM")
	msg        = []byte("email message...")
	recipients = []string{"tymalik@pm.me"}
)

func main() {
	hostname := os.Getenv("HOSTNAME")
	auth := smtp.PlainAuth("", username, password, hostname)

	err := smtp.SendMail(hostname+":587", auth, from, recipients, msg)
	if err != nil {
		log.Fatal(err)
	}

	// c, err := smtp.Dial("smtp.gmail.com:587")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := c.Mail("ty.p.malik@gmail.com"); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := c.Rcpt("tymalik@pm.me"); err != nil {
	// 	log.Fatal(err)
	// }

	// wc, err := c.Data()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = fmt.Fprintf(wc, "Email body...")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = wc.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = c.Quit()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
