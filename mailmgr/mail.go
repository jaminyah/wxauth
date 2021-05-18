package mailmgr

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Send(emailAddr string, authCode string) int {

	components := strings.Split(emailAddr, "@")

	from := mail.NewEmail("WxAlert Sign-up", "noreply@wxalert.us")
	subject := "Activation Code"
	to := mail.NewEmail(components[0], emailAddr)
	plainTextContent := "Your activation code is: " + authCode
	htmlContent := "<strong>Use the code to activate your WxAlert account: </strong>" + authCode
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	mailServer := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := mailServer.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return response.StatusCode
}
