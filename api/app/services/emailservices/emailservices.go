package emailservices

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var smtp string
var user string
var password string
var tlsPort int

type EmailStruct struct {
	From    string
	To      string
	Cc      string
	Subject string
	Body    string
	Attach  string
}

func init() {
	smtp = os.Getenv("GMail_SMTP")
	user = os.Getenv("GMail_User")
	password = os.Getenv("GMail_Password")
	tlsPort, _ = strconv.Atoi(os.Getenv("GMail_TLS_Port"))
}

func Send(es EmailStruct) {
	if es.From == "" {
		es.From = "kimi0x03@gmail.com"
	}

	if es.Body == "" {
		es.Body = "Hello kimi!"
	}

	if es.Subject == "" {
		es.Subject = "Send From Kimi0x03"
	}
	mail := gomail.NewMessage()
	mail.SetHeader("From", es.From)
	mail.SetHeader("To", es.To)
	// m.SetAddressHeader("Cc", "kimi.tsai@eten.com.tw", "")
	mail.SetHeader("Subject", es.Subject)
	mail.SetBody("text/html", es.Body)
	// m.Attach("/home/Alex/lolcat.jpg")

	dialer := gomail.NewDialer(smtp, tlsPort, user, password)

	// Send the email
	if err := dialer.DialAndSend(mail); err != nil {
		// panic(err)
		fmt.Println("---------------- emailService:", err)
	}
}
