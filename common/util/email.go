package util

import (
	"github.com/astaxie/beego"
	"net/smtp"
	"strings"
)

func SendMail(to, subject, body, mailType string) error {
	var (
		user     = beego.AppConfig.String("emailuser")
		password = beego.AppConfig.String("emailpassword")
		host     = beego.AppConfig.String("emailhost")
		port     = beego.AppConfig.String("emailport")
	)
	auth := smtp.PlainAuth("", user, password, host)

	var contentType string
	if EqualString(mailType, "html") {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host+":"+port, auth, user, sendTo, msg)

	return err
}
