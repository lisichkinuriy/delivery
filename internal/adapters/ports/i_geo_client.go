package ports

import (
	"context"
	"lisichkinuriy/delivery/internal/domain/vo"
)

type IGeoClient interface {
	GetLocation(ctx context.Context, street string) (vo.Location, error)
}
