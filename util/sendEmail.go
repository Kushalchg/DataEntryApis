package util

import (
	"net/smtp"

	"www.github.com/kushalchg/DataEntryApis/global"
)

func SendMail() {
	from := "kushalchapagain123456@gmail.com"
	password := "aiou fkok djyl jdqb"
	// password := "#Learn1gmail"
	msg := []byte("this is example message and your code is 330232")
	to := []string{"kushalchapagain74@gmail.com"}
	host := "smtp.gmail.com"
	auth := smtp.PlainAuth("this is identify", from, password, host)
	err := smtp.SendMail(host+":587", auth, from, to, msg)
	if err != nil {
		global.Logger.Print("error while sending message", err)
	} else {
		global.Logger.Print("email sent successfully")
	}
}
