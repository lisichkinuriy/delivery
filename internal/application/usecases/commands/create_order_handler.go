package commands

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/domain/order"
)

type ICreateOrderHandler interface {
	Handle(context.Context, CreateOrderCommand) error
}

var _ ICreateOrderHandler = &CreateOrderHandler{}

type CreateOrderHandler struct {
	orderRepo ports.IOrderRepository
	geoClient ports.IGeoClient
}

func NewCreateOrderHandler(orderRepo ports.IOrderRepository, geoClient ports.IGeoClient) (*CreateOrderHandler, error) {
	if orderRepo == nil {
		return nil, errors.New("order repository is nil")
	}
	if geoClient == nil {
		return nil, errors.New("geo client is nil")
	}
	return &CreateOrderHandler{orderRepo: orderRepo, geoClient: geoClient}, nil
}

func (ch *CreateOrderHandler) Handle(ctx context.Context, command CreateOrderCommand) error {
	if command.isEmpty() {
		return errors.New("invalid command")
	}

	orderAggregate, err := ch.orderRepo.Get(ctx, command.OrderId())
	if orderAggregate != nil {
		return errors.New("order already exists")
	}

	location, err := ch.geoClient.GetLocation(ctx, command.Street())
	if err != nil {
		return err
	}
	log.Info(location.X(), location.Y())

	newOrder, err := order.NewOrder(command.OrderId(), location)
	if err != nil {
		return err
	}

	err = ch.orderRepo.Add(ctx, newOrder)
	if err != nil {
		return err
	}

	return nil
}
