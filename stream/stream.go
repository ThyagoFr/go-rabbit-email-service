package stream

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
	"github.com/thyagofr/go-rabbit-email-service/configs"
)

func InitStream(config *configs.RabbitMQConfig) {

	conn, err := amqp.Dial(config.URL())
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	msgsChannel, err := ch.Consume(
		config.QUEUE,
		config.CONSUMER,
		true,
		false,
		false,
		false,
		nil,
	)
	loop := make(chan bool)
	go goListenQueue(msgsChannel)
	<-loop

}

func goListenQueue(messages <-chan amqp.Delivery) {
	for msg := range messages {
		fmt.Printf("Mensagem teste %s \n", msg.Body)
	}
}
