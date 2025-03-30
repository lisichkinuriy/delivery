package commands

import (
	"context"
	"errors"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/domain/order"
	"lisichkinuriy/delivery/internal/domain/vo"
)

type ICreateOrderHandler interface {
	Handle(context.Context, CreateOrderCommand) error
}

var _ ICreateOrderHandler = &CreateOrderHandler{}

type CreateOrderHandler struct {
	orderRepo ports.IOrderRepository
}

func NewCreateOrderHandler(orderRepo ports.IOrderRepository) (*CreateOrderHandler, error) {
	if orderRepo == nil {
		return nil, errors.New("order repository is nil")
	}

	return &CreateOrderHandler{orderRepo: orderRepo}, nil
}

func (ch *CreateOrderHandler) Handle(ctx context.Context, command CreateOrderCommand) error {
	if command.isEmpty() {
		return errors.New("invalid command")
	}

	orderAggregate, err := ch.orderRepo.Get(ctx, command.OrderId())
	if orderAggregate != nil {
		return errors.New("order already exists")
	}

	location, err := vo.FakeLocation()
	if err != nil {
		return err
	}

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
