package email

import (
	"blog-go/global"
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

// Email sends an email to one or multiple recipients using gomail
func Email(To, subject, body string) error {
	emailCfg := global.Config.Email // load email settings from global config

	from := emailCfg.From
	// nickname := emailCfg.Nickname
	secret := emailCfg.Secret
	host := emailCfg.Host
	port := emailCfg.Port
	isSSL := emailCfg.IsSSL

	// create a new gomail message
	m := gomail.NewMessage()

	// set sender with optional nickname
	// if nickname != "" {
	// 	m.SetHeader("From", m.FormatAddress(from, nickname))
	// } else {
	// 	m.SetHeader("From", from)
	// }
	m.SetHeader("From", from)

	// split recipients by comma and set
	m.SetHeader("To", To)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// create dialer
	d := gomail.NewDialer(host, port, from, secret)

	d.SSL = isSSL

	// send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
