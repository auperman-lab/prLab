package manager

import (
	"bytes"
	"encoding/json"
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
)

type Consumer struct {
	url       string
	ch        *amqp.Channel
	queueName string
}

func NewConsumer(url string, ch *amqp.Channel, queueName string) *Consumer {
	return &Consumer{
		url:       url,
		ch:        ch,
		queueName: queueName,
	}
}

func (c *Consumer) Consume() {

	q, err := c.ch.QueueDeclare(
		c.queueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Panicf("Failed to declare a queue: %s", err)
	}

	msgs, err := c.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Panicf("Failed to publish a message: %s", err)
	}
	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			product, err := deserializeProduct(d.Body)
			if err != nil {
				log.Printf("Failed to deserialize a product: %s", err)
			} else {
				log.Printf("Valid Product: %+v", product)
				sendRequest(product, c.url)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

func deserializeProduct(data []byte) (Product, error) {
	var product Product
	if err := json.Unmarshal(data, &product); err != nil {
		log.Printf("Error unmarshalling message: %s", err)
		return Product{}, err
	}
	if product.Name == "" || product.Price <= 0 {
		log.Printf("Invalid product data: %+v", product)
		return Product{}, errors.New("invalid product data")
	}

	return product, nil
}

func sendRequest(product Product, backendURL string) {
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Printf("Error serializing product to JSON: %s", err)
		return
	}

	resp, err := http.Post(backendURL, "application/json", bytes.NewBuffer(productJSON))
	if err != nil {
		log.Printf("Failed to send POST request: %s", err)
		return

	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Backend responded with status: %d", resp.StatusCode)
	} else {
		log.Printf("Successfully sent product to backend.")
	}
	return
}

func (c *Consumer) UpdateURL(newURL string) {
	c.url = newURL + "/products"
}
