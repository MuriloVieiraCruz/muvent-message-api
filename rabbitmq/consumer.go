package rabbitmq

import (
	"encoding/json"
	"log"
	//"time"

	"github.com/streadway/amqp"
	//"github.com/gocql/gocql"
	"muvent-message-api/models"
	"muvent-message-api/email"
)

func ConsumeMessages(conn *amqp.Connection) {

	/*cluster := gocql.NewCluster("127.0.0.1") // Altere para o IP do seu Cassandra
    cluster.Keyspace = "your_keyspace"       // Substitua pelo nome do seu keyspace
    cluster.Consistency = gocql.Quorum
    cluster.Timeout = 10 * time.Second

	session, err := cluster.CreateSession()
    if err != nil {
        log.Fatalf("Failed to connect to Cassandra: %v", err)
    }
    defer session.Close()*/

	channel, err := conn.Channel()

	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"send-email-queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var user models.EmailRequest

			err := json.Unmarshal(d.Body, &user)

			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				continue
			}

			log.Printf("Received a message: %+v", user)
			email.SendEmail(user)

			/*if err := save(session, &user); err != nil {
				log.Printf("Failed to save data to Cassandra: %v", err)
                continue
			}*/
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<- forever
}

/*func save(session *gocql.Session, user *models.EmailRequest) error {
    query := `INSERT INTO email_requests (id, email, firstName, lastName, timestamp) VALUES (?, ?, ?, ?, ?)`
    return session.Query(query, gocql.TimeUUID(), user.Email, user.FirstName, user.LastName, user.Timestamp).Exec()
}*/

