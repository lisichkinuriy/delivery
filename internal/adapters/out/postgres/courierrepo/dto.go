package courierrepo

import (
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/adapters/out/postgres/shared"
	"lisichkinuriy/delivery/internal/domain/courier"
)

type CourierDTO struct {
	ID        uuid.UUID          `gorm:"type:uuid;primaryKey"`
	Name      string             `gorm:"type:varchar(255);not null"`
	Transport TransportDTO       `gorm:"foreignKey:CourierID;constraint:OnDelete:CASCADE;"`
	Location  shared.LocationDTO `gorm:"embedded;embeddedPrefix:location_"`
	Status    courier.Status     `gorm:"type:varchar(255);not null"`
}

type TransportDTO struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Speed     int       `gorm:"type:int;not null"`
	CourierID uuid.UUID `gorm:"type:uuid;index"`
}

func (CourierDTO) TableName() string {
	return "couriers"
}

func (TransportDTO) TableName() string {
	return "transports"
}
