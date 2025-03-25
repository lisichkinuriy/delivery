package courierrepo

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"lisichkinuriy/delivery/internal/adapters/out/postgres"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/domain/courier"
)

var _ ports.ICourierRepository = &Repository{}

type Repository struct {
	db *gorm.DB
}

func (r Repository) Add(ctx context.Context, aggregate *courier.Courier) error {
	dto := DomainToDTO(aggregate)

	tx := postgres.GetTxFromContext(ctx)
	if tx == nil {
		tx = r.db
	}
	err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Create(&dto).Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) Update(ctx context.Context, aggregate *courier.Courier) error {
	dto := DomainToDTO(aggregate)

	tx := postgres.GetTxFromContext(ctx)
	if tx == nil {
		tx = r.db
	}
	err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Save(&dto).Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) Get(ctx context.Context, id uuid.UUID) (*courier.Courier, error) {
	dto := CourierDTO{}

	tx := postgres.GetTxFromContext(ctx)
	if tx == nil {
		tx = r.db
	}
	result := tx.
		Preload(clause.Associations).
		Find(&dto, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	aggregate := DtoToDomain(dto)
	return aggregate, nil
}

func (r Repository) GetAllFreeCouriers(ctx context.Context) ([]*courier.Courier, error) {
	var dtos []CourierDTO

	tx := postgres.GetTxFromContext(ctx)
	if tx == nil {
		tx = r.db
	}

	result := tx.Preload(clause.Associations).
		Where("status = ?", courier.StatusFree).
		Find(&dtos)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	aggregates := make([]*courier.Courier, len(dtos))
	for i, dto := range dtos {
		aggregates[i] = DtoToDomain(dto)
	}

	return aggregates, nil
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	if db == nil {
		return nil, errors.New("db connection is nil")
	}

	return &Repository{
		db: db,
	}, nil
}
