package util

import (
	"net/smtp"
	"os"
)

func SendMail(to []string, msg []byte) error {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EPASS")
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("this is identify", from, password, host)
	err := smtp.SendMail(host+":587", auth, from, to, msg)

	return err
}
