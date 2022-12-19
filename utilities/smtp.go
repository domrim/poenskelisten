package utilities

import (
	"aunefyren/poenskelisten/config"
	"aunefyren/poenskelisten/models"
	"log"

	"github.com/go-mail/mail"
)

func SendSMTPVerificationEmail(user models.User) error {

	// Get configuration
	config, err := config.GetConfig()
	if err != nil {
		return err
	}

	log.Println("Sending e-mail to user " + user.FirstName + " " + user.LastName + ".")

	m := mail.NewMessage()
	m.SetHeader("From", "Pønskelisten <"+config.SMTPFrom+">")
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Please verify your account")
	m.SetBody("text/html", "Hello <b> + user.FirstName + </b>!<br><br>This is a test alert saying to verify your account.")

	d := mail.NewDialer(config.SMTPHost, config.SMTPPort, config.SMTPUsername, config.SMTPPassword)

	// Send the email
	err = d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
