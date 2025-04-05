package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/application/usecases/commands"
	"lisichkinuriy/delivery/pkg/queues/basketconfirmedpb/queues/basketconfirmedpb"
	"log"
)

var _ ports.IBasketConfirmedConsumer = &BasketConfirmedConsumer{}

type BasketConfirmedConsumer struct {
	topic                     string
	consumer                  *kafka.Consumer
	createOrderCommandHandler *commands.CreateOrderHandler
}

func (c *BasketConfirmedConsumer) Consume() error {
	err := c.consumer.Subscribe(c.topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	for {
		ctx := context.Background()
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

		// Обрабатываем сообщение
		fmt.Printf("Received: %s => %s\n", msg.TopicPartition, string(msg.Value))

		var event basketconfirmedpb.BasketConfirmedIntegrationEvent
		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
		}

		// Отправляем команду
		createOrderCommand, err := commands.NewCreateOrderCommand(
			uuid.MustParse(event.BasketId), event.Address.Street)
		if err != nil {
			log.Printf("Failed to create createOrder command: %v", err)
		} else {
			err = c.createOrderCommandHandler.Handle(ctx, createOrderCommand)
			if err != nil {
				log.Printf("Failed to handle createOrder command: %v", err)
			}
		}

		// Подтверждаем обработку сообщения
		_, err = c.consumer.CommitMessage(msg)
		if err != nil {
			log.Printf("Commit failed: %v", err)
		}
	}

}

func (b *BasketConfirmedConsumer) Close() error {
	return b.consumer.Close()
}

func NewBasketConfirmedConsumer(createOrderCommandHandler *commands.CreateOrderHandler) (*BasketConfirmedConsumer, error) {
	if createOrderCommandHandler == nil {
		return nil, errors.New("create order handler is nil")
	}

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"group.id":           "delivery-service-group",
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}

	return &BasketConfirmedConsumer{
		topic:                     "basket.confirmed",
		consumer:                  consumer,
		createOrderCommandHandler: createOrderCommandHandler,
	}, nil
}
