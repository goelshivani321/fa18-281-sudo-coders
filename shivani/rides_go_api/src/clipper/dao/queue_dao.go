package dao

import (
	"github.com/streadway/amqp"
	"log"
)

type QueueDAO struct {
	User   string
	Password string
	Server string
	Port string
}

var quChan *amqp.Channel

// Establish a connection to rabbitmq
func (qu *QueueDAO) Connect() {
	log.Println(qu.User)
	log.Println(qu.Password)
	conn, err := amqp.Dial("amqp://"+qu.User+":"+qu.Password+"@"+qu.Server+":"+qu.Port+"/")
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	quChan= ch
}

// Send ride to Queue for Processing
func (qu *QueueDAO) QueueSend(message string, queue string) error {

	q, err := quChan.QueueDeclare(
		queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return err
	}

	body := message
	err = quChan.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	return err
}

// Receive ride from Queue to Process
func (qu *QueueDAO) QueueReceive(queue string) ([]string, error) {

	q, err := quChan.QueueDeclare(
		queue, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return nil, err
	}

	msgs, err := quChan.Consume(
		q.Name, // queue
		"rides",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return nil, err
	}

	rideIds := make(chan string)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			rideIds <- string(d.Body)
		}
		close(rideIds)
	}()

	err = quChan.Cancel("rides", false)

	var rideIdsArray []string
	for n := range rideIds {
		rideIdsArray = append(rideIdsArray, n)
	}

	return rideIdsArray, err
}