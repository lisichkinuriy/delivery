package ports

import (
	"context"
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/domain/courier"
)

type ICourierRepository interface {
	Add(ctx context.Context, aggregate *courier.Courier) error
	Update(ctx context.Context, aggregate *courier.Courier) error
	Get(ctx context.Context, id uuid.UUID) (*courier.Courier, error)
	GetAllFreeCouriers(ctx context.Context) ([]*courier.Courier, error)
}
