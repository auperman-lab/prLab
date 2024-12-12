package main

import (
	"github.com/auperman-lab/lab3/manager"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	defer ch.Close()

	url := "http://localhost:9001/products"
	queName := "products"
	consumer := manager.NewConsumer(url, ch, queName)
	go consumer.Consume()

	ftpClient := manager.NewFTPClient("ftp_server:21", "testuser", "testpass", "http://node1:9001")
	go ftpClient.PollAndProcess()

	listener := manager.NewListener("8080", consumer, ftpClient)

	err = listener.Start()
	if err != nil {
		panic(err)

	}

	select {}
}
