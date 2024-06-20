package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net"
	"net/smtp"

	"github.com/pkg/errors"
)

type Client struct {
	Host string
	Port int
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

func Send(from, receiverEmail string, auth smtp.Auth, cli *Client) {

	// Receiver email address.
	to := []string{
		receiverEmail,
	}

	// smtp server configuration.
	smtpHost := cli.Host
	smtpPort := cli.Port

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", smtpHost, smtpPort))
	if err != nil {
		println(err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		println(err)
	}

	tlsConfig := &tls.Config{
		ServerName: smtpHost,
	}

	if err = c.StartTLS(tlsConfig); err != nil {
		println(err)
	}

	if err := c.Auth(auth); err != nil {
		println(err)
		return
	}

	t, err := template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	_ = t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "Hasan Yousef",
		Message: "This is a test message in a HTML template",
	})

	// Sending email.
	err = smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
