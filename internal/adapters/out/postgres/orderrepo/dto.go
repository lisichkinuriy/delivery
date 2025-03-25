package orderrepo

import (
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/shared"
	"lisichkinuriy/delivery/internal/domain/order"
)

type OrderDTO struct {
	ID        uuid.UUID          `gorm:"type:uuid;primaryKey"`
	Location  shared.LocationDTO `gorm:"embedded;embeddedPrefix:location_"`
	Status    order.Status       `gorm:"type:varchar(20)"`
	CourierID *uuid.UUID         `gorm:"type:uuid;index"`
}

func (OrderDTO) TableName() string {
	return "orders"
}
