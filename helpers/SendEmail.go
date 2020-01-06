package helpers

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

/*
 *	Function to send otp to user email address
 *
 *	return Response
 */
func SendOtpMail(fullName, toEmail string, OTP string) (err error) {
	from := mail.NewEmail("Test Institute", "test@example.com")
	subject := "Comfirmation Mail"
	to := mail.NewEmail(fullName, toEmail)
	// plainTextContent := ""
	htmlContent := "Your confirmation OTP is = " + OTP
	message := mail.NewSingleEmail(from, subject, to, " ", htmlContent) // ("" = plainTextContent)
	client := sendgrid.NewSendClient("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")
	_, err = client.Send(message)
	return err
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(response.StatusCode)
	// fmt.Println(response.Body)
	// fmt.Println(response.Headers)
	// }
}
