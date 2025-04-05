package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/domain/order"
	"lisichkinuriy/delivery/pkg/queues/orderstatuschangedpb/queues/orderstatuschangedpb"
	"log"
)

type OrderProducer struct {
	sarama sarama.AsyncProducer
}

var _ ports.IOrderProducer = &OrderProducer{}

func NewOrderProducer() (*OrderProducer, error) {
	version, err := sarama.ParseKafkaVersion("3.4.0")
	if err != nil {
		return nil, fmt.Errorf("parse Kafka version: %w", err)
	}

	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.RequiredAcks = sarama.WaitForAll
	saramaCfg.Producer.Retry.Max = 5
	saramaCfg.Producer.Return.Successes = true
	saramaCfg.Producer.Return.Errors = true
	saramaCfg.Producer.Partitioner = sarama.NewHashPartitioner
	saramaCfg.Version = version

	addrs := []string{"localhost:9092"}
	producer, err := sarama.NewAsyncProducer(addrs, saramaCfg)
	if err != nil {
		return nil, fmt.Errorf("create async producer: %w", err)
	}

	// обработка ошибок/успехов
	go func() {
		for {
			select {
			case msg := <-producer.Successes():
				log.Printf("Успешно: partition=%d offset=%d", msg.Partition, msg.Offset)
			case err := <-producer.Errors():
				log.Printf("Ошибка: %v", err)
			}
		}
	}()

	return &OrderProducer{
		sarama: producer,
	}, nil
}

func (p *OrderProducer) Publish(ctx context.Context, domainEvent order.CompletedDomainEvent) error {
	integrationEvent, err := p.mapDomainEventToIntegrationEvent(domainEvent)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(integrationEvent)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "order.status.changed",
		Key:   sarama.StringEncoder(domainEvent.ID().String()),
		Value: sarama.ByteEncoder(bytes),
	}

	select {
	case p.sarama.Input() <- msg:
		return nil
	case <-ctx.Done():
		return fmt.Errorf("send cancelled: %w", ctx.Err())
	}
}

func (p *OrderProducer) Close() error {
	return p.sarama.Close()
}

func (p *OrderProducer) mapDomainEventToIntegrationEvent(domainEvent order.CompletedDomainEvent) (*orderstatuschangedpb.OrderStatusChangedIntegrationEvent, error) {
	status, ok := orderstatuschangedpb.OrderStatus_value[domainEvent.OrderStatus()]
	if !ok {
		return nil, errors.New("order status not found")
	}

	integrationEvent := orderstatuschangedpb.OrderStatusChangedIntegrationEvent{
		OrderId:     domainEvent.OrderID().String(),
		OrderStatus: orderstatuschangedpb.OrderStatus(status),
	}
	return &integrationEvent, nil
}
