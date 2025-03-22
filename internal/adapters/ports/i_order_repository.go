package ports

import (
	"context"
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/domain/order"
)

type IOrderRepository interface {
	Add(ctx context.Context, aggregate *order.Order) error
	Update(ctx context.Context, aggregate *order.Order) error
	Get(ctx context.Context, id uuid.UUID) (*order.Order, error)
	GetFirstCreatedOrder(ctx context.Context) (*order.Order, error)
	GetAssignedOrders(ctx context.Context) ([]*order.Order, error)
}
