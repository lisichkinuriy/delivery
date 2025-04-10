package pkg

import (
	"github.com/google/uuid"
)

type IDomainEvent interface {
	GetEventID() uuid.UUID
	GetEventName() string
}
