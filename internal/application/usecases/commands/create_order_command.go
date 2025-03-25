package commands

import (
	"errors"
	"github.com/google/uuid"
)

type CreateOrderCommand struct {
	orderId uuid.UUID
	street  string
	isSet   bool
}

func NewCreateOrderCommand(orderId uuid.UUID, street string) (CreateOrderCommand, error) {
	if orderId == uuid.Nil {
		return CreateOrderCommand{}, errors.New("orderId is required")
	}

	if street == "" {
		return CreateOrderCommand{}, errors.New("street is required")
	}

	return CreateOrderCommand{orderId: orderId, street: street, isSet: true}, nil
}

func (c CreateOrderCommand) isEmpty() bool {
	return !c.isSet
}

func (c CreateOrderCommand) OrderId() uuid.UUID { return c.orderId }
func (c CreateOrderCommand) Street() string     { return c.street }
