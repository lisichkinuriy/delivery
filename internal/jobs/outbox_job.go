package jobs

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/robfig/cron/v3"
	"lisichkinuriy/delivery/internal/adapters/outbox"
	"lisichkinuriy/delivery/internal/domain/order"
	"time"
)

var _ cron.Job = &OutboxJob{}

type OutboxJob struct {
	outboxRepository outbox.IOutboxRepository
	eventRegistry    outbox.IEventRegistry
}

func NewOutboxJob(outboxRepository outbox.IOutboxRepository, eventRegistry outbox.IEventRegistry) (*OutboxJob, error) {
	if outboxRepository == nil {
		return nil, errors.New("outbox repository is nil")
	}
	if eventRegistry == nil {
		return nil, errors.New("event registry is nil")
	}

	return &OutboxJob{
		outboxRepository: outboxRepository,
		eventRegistry:    eventRegistry}, nil
}

func (j *OutboxJob) Run() {
	ctx := context.Background()

	log.Info("starting outbox job")
	// Получаем не отправленные Outbox Events
	outboxMessages, err := j.outboxRepository.GetNotPublishedMessages()
	if err != nil {
		log.Error(err)
	}

	// Перебираем в цикле
	for _, outboxMessage := range outboxMessages {
		// Приводим Outbox Message -> Domain Event
		domainEvent, err := j.eventRegistry.DecodeDomainEvent(outboxMessage)
		if err != nil {
			log.Error(err)
			continue
		}

		// Go не поддерживает вызов generic-функций с параметрами T во время выполнения
		// Поэтому делаем Switch
		switch domainEvent.(type) {
		case *order.CompletedDomainEvent:
			err := mediatr.Publish[*order.CompletedDomainEvent](ctx,
				domainEvent.(*order.CompletedDomainEvent))
			if err != nil {
				log.Error(err)
				continue
			}
		}

		// Если ошибок нет, помечаем Outbox Message как отправленное и сохраняем в БД
		// А если были ошибки, то цикл просто повторяется
		now := time.Now().UTC()
		outboxMessage.ProcessedOnUtc = &now
		err = j.outboxRepository.Update(ctx, outboxMessage)
		if err != nil {
			log.Error(err)
			continue
		}
	}
}
