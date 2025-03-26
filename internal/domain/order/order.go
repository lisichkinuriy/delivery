package order

import (
	"errors"
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/domain/courier"
	"lisichkinuriy/delivery/internal/domain/vo"
)

type Order struct {
	id        uuid.UUID
	location  vo.Location
	status    Status
	courierID *uuid.UUID
}

func (o *Order) Status() Status        { return o.status }
func (o *Order) CourierID() *uuid.UUID { return o.courierID }
func (o *Order) Location() vo.Location { return o.location }
func (o *Order) ID() uuid.UUID         { return o.id }

func (o *Order) IsCreated() bool { return o.Status() == StatusCreated }

func NewOrder(id uuid.UUID, l vo.Location) (*Order, error) {
	if id == uuid.Nil {
		return nil, errors.New("id should not be nil")
	}
	if l.IsEmpty() {
		return nil, errors.New("empty location")
	}

	return &Order{
		id:        id,
		location:  l,
		status:    StatusCreated,
		courierID: nil,
	}, nil
}

func (o *Order) Assign(courier *courier.Courier) error {

	if courier == nil {
		return errors.New("courier should not be nil")
	}

	if courier.IsBusy() {
		return errors.New("courier is busy")
	}

	if o.status != StatusCreated {
		return errors.New("wrong status")
	}

	o.status = StatusAssigned
	courierID := courier.ID()
	o.courierID = &courierID

	return nil
}

func (o *Order) Complete() error {
	if o.status != StatusAssigned {
		return errors.New("order status is not assigned")
	}
	o.status = StatusCompleted
	return nil
}

func RestoreOrder(id uuid.UUID, courierID *uuid.UUID, location vo.Location, status Status) *Order {
	return &Order{
		id:        id,
		courierID: courierID,
		location:  location,
		status:    status,
	}
}
