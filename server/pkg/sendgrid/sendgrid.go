package sgclient

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	ToName  string
	ToEmail string

	FromName  string
	FromEmail string

	Subject  string
	BodyText string
	BodyHTML string
}

type SendgridClient struct {
	apiKey string
	Client *sendgrid.Client
}

func NewSendgridClient() *SendgridClient {
	ret := new(SendgridClient)

	ret.apiKey = os.Getenv("SENDGRID_API_KEY")
	ret.Client = sendgrid.NewSendClient(ret.apiKey)

	return ret
}

func (s *SendgridClient) SendEmail(email Email) error {
	from := mail.NewEmail(email.FromName, email.FromEmail)
	to := mail.NewEmail(email.ToName, email.ToEmail)
	message := mail.NewSingleEmail(
		from, email.Subject, to, email.BodyText, email.BodyHTML,
	)

	response, err := s.Client.Send(message)
	if err != nil {
		return err
	}
	if response.StatusCode != 202 {
		fmt.Printf("Unknown status code %d\n\t%+v\n\t%s\n", response.StatusCode, response.Headers, response.Body)
		return fmt.Errorf("unknown status code %d is not 202", response.StatusCode)
	}

	return nil
}
