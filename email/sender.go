package email

import (
	"html/template"
	"bytes"
    "fmt"
    "log"
    "os"

	"muvent-message-api/models"
    "github.com/joho/godotenv"
    "github.com/sendgrid/sendgrid-go"
    "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(user models.EmailRequest) {

    err := godotenv.Load("sendgrid.env")
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    tmpl, err := template.ParseFiles("./template/email_template.html")
    if err != nil {
        log.Printf("Error parsing template: %s", err)
        return
    }

    var body bytes.Buffer
    if err := tmpl.Execute(&body, user); err != nil {
        log.Printf("Error executing template: %s", err)
        return
    }

    // Configurar o e-mail
    fromEmail := mail.NewEmail("Example User", "murilo12super@gmail.com")
    toEmail := mail.NewEmail("Recipient Name", user.Email)
    subject := fmt.Sprintf("Welcome, %s!", user.FirstName)
    content := mail.NewContent("text/html", body.String())

    m := mail.NewV3MailInit(fromEmail, subject, toEmail, content)

    // Enviar o e-mail usando SendGrid
    client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
    response, err := client.Send(m)
    if err != nil {
        log.Printf("Error sending email: %s", err)
        return
    }

    log.Printf("Email sent successfully to %s", user.Email)
    log.Printf("Response Status Code: %d", response.StatusCode)
    log.Printf("Response Body: %s", response.Body)
    log.Printf("Response Headers: %v", response.Headers)

}