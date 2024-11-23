package mail

import (
	"context"
	"net/smtp"
)

type connection struct {
	host     string
	port     string
	username string
	password string
}

func New(host, port, username, password string) *connection {
	return &connection{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (c *connection) Prepare() *email {
	return &email{}
}

func (c *connection) Send(ctx context.Context, email *email) error {
	auth := smtp.PlainAuth("", c.username, c.password, c.host)
	smtpAddress := c.host + ":" + c.port

	err := smtp.SendMail(smtpAddress, auth, c.username, append(email.to, email.cc...), email.buildBytes())
	if err != nil {
		return err
	}

	return nil
}
