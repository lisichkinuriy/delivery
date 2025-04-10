package order

import (
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/domain/pkg"
)

const CompletedDomainEventName = "CompletedDomainEvent"

var _ pkg.IDomainEvent = CompletedDomainEvent{}

type CompletedDomainEvent struct {
	// base
	Id   uuid.UUID
	Name string

	// payload
	OrderID     uuid.UUID
	OrderStatus string

	isSet bool
}

func (e CompletedDomainEvent) GetEventName() string {
	return e.Name
}

func (e CompletedDomainEvent) GetEventID() uuid.UUID { return e.Id }

func NewCompletedDomainEvent(aggregate *Order) CompletedDomainEvent {
	return CompletedDomainEvent{
		Id:   uuid.New(),
		Name: CompletedDomainEventName,

		OrderID:     aggregate.ID(),
		OrderStatus: aggregate.Status().String(),

		isSet: true,
	}
}

func (e CompletedDomainEvent) IsEmpty() bool {
	return !e.isSet
}
