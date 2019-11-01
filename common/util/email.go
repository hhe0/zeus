package util

import (
	"github.com/astaxie/beego"
	"net/smtp"
	"strings"
)

type Mail struct {
	To       string `json:"to"`
	Subject  string `json:"string"`
	Body     string `json:"body"`
	MailType string `json:"mailType"`
}

func SendMail(mail Mail) error {
	var (
		user     = beego.AppConfig.String("emailuser")
		password = beego.AppConfig.String("emailpassword")
		host     = beego.AppConfig.String("emailhost")
		port     = beego.AppConfig.String("emailport")
	)
	auth := smtp.PlainAuth("", user, password, host)

	var contentType string
	if EqualString(mail.MailType, "html") {
		contentType = "Content-Type: text/" + mail.MailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + mail.To + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + mail.Subject + "\r\n" + contentType + "\r\n\r\n" + mail.Body)
	sendTo := strings.Split(mail.To, ";")
	err := smtp.SendMail(host+":"+port, auth, user, sendTo, msg)

	return err
}
