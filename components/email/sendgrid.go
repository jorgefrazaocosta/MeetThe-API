package email

import (
	"fmt"
	"log"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailModel struct {
	FromName    string
	FromEmail   string
	ToName      string
	ToEmail     string
	Subject     string
	Message     string
	MessageHTML string
}

func SendEmail(model EmailModel) bool {

	from := mail.NewEmail(model.FromName, model.FromEmail)
	to := mail.NewEmail(model.ToName, model.ToEmail)
	plainTextContent := model.Message
	htmlContent := model.MessageHTML
	message := mail.NewSingleEmail(from, model.Subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {

		log.Println(err)
		return false

	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return true

}
