package mailer

const (
	FromName   = "GopherSocial"
	maxRetires = 3
)

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) error
}
