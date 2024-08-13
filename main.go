package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {

	err := godotenv.Load("sendgrid.env")
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    from := mail.NewEmail("Example User", "murilo12super@gmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "murilo12super@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

/*package main

import (
    "log"
	"muvent-message-api/rabbitmq"
)

func main() {
    conn, err := rabbitmq.ConnectRabbitMQ()
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    rabbitmq.ConsumeMessages(conn)
}*/