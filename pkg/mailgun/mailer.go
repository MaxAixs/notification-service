package mailgun

type Mailer interface {
	SendEmail(to, topic, body string) error
}
