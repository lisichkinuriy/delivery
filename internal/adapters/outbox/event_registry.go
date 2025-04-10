package outbox

import (
	"encoding/json"
	"errors"
	"fmt"
	"lisichkinuriy/delivery/internal/domain/pkg"
	"reflect"
	"time"
)

type IEventRegistry interface {
	DecodeDomainEvent(event *Message) (pkg.IDomainEvent, error)
}

var _ IEventRegistry = &EventRegistry{}

type EventRegistry struct {
	EventRegistry map[string]reflect.Type
}

func NewEventRegistry() (*EventRegistry, error) {
	return &EventRegistry{
		EventRegistry: make(map[string]reflect.Type),
	}, nil
}

func (r *EventRegistry) RegisterDomainEvent(eventType reflect.Type) error {
	if eventType == nil {
		return errors.New("eventType is nil")
	}
	eventName := eventType.Name()
	r.EventRegistry[eventName] = eventType
	return nil
}

func EncodeDomainEvent(domainEvent pkg.IDomainEvent) (Message, error) {
	payload, err := json.Marshal(domainEvent)
	if err != nil {
		return Message{}, fmt.Errorf("failed to marshal event: %w", err)
	}

	return Message{
		ID:             domainEvent.GetEventID(),
		Name:           domainEvent.GetEventName(),
		Payload:        payload,
		OccurredOnUtc:  time.Now().UTC(),
		ProcessedOnUtc: nil,
	}, nil
}

func EncodeDomainEvents(domainEvent []pkg.IDomainEvent) ([]Message, error) {
	outboxMessages := make([]Message, 0)
	for _, event := range domainEvent {
		event, err := EncodeDomainEvent(event)
		if err != nil {
			return nil, err
		}
		outboxMessages = append(outboxMessages, event)
	}
	return outboxMessages, nil
}

func (r *EventRegistry) DecodeDomainEvent(outboxMessage *Message) (pkg.IDomainEvent, error) {
	t, ok := r.EventRegistry[outboxMessage.Name]
	if !ok {
		return nil, fmt.Errorf("unknown outboxMessage type: %s", outboxMessage.Name)
	}

	// Создаём новый указатель на нужный тип
	eventPtr := reflect.New(t).Interface()

	if err := json.Unmarshal(outboxMessage.Payload, eventPtr); err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	// Приводим к DomainEvent
	domainEvent, ok := eventPtr.(pkg.IDomainEvent)
	if !ok {
		return nil, fmt.Errorf("decoded outboxMessage does not implement DomainEvent")
	}

	return domainEvent, nil
}
