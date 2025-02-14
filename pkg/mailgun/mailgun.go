package mailgun

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
	"time"
)

type MgClient struct {
	mg     *mailgun.MailgunImpl
	domain string
}

func NewMailgun(domain, apiKey string) *MgClient {
	mg := mailgun.NewMailgun(domain, apiKey)
	return &MgClient{mg: mg, domain: domain}
}

func (m *MgClient) SendEmail(to, topic, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	from := fmt.Sprintf("postmaster@%s", m.domain)

	logrus.Infof("Sending email from %s to %s with topic: %s, body: %s", from, to, topic, body)
	message := mailgun.NewMessage(from, topic, body, to)

	_, id, err := m.mg.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("cannot send msg to email: %w", err)
	}

	logrus.Infof("Email sent successfully. Message ID: %s", id)
	return nil
}
