package ports

import (
	"context"
	"lisichkinuriy/delivery/internal/domain/order"
)

type IOrderProducer interface {
	Publish(ctx context.Context, domainEvent order.CompletedDomainEvent) error
	Close() error
}
