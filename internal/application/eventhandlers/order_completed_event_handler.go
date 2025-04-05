package eventhandlers

import (
	"context"
	"errors"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/domain/order"
)

type OrderCompletedDomainEventHandler struct {
	orderProducer ports.IOrderProducer
}

func NewOrderCompletedDomainEventHandler(orderProducer ports.IOrderProducer) (*OrderCompletedDomainEventHandler, error) {

	if orderProducer == nil {
		return nil, errors.New("orderProducer is nil!")
	}

	return &OrderCompletedDomainEventHandler{orderProducer: orderProducer}, nil
}

func (o *OrderCompletedDomainEventHandler) Handle(ctx context.Context, notification order.CompletedDomainEvent) error {
	return o.orderProducer.Publish(ctx, notification)
}

var _ IEventHandler[order.CompletedDomainEvent] = &OrderCompletedDomainEventHandler{}
