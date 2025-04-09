package outbox

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type IOutboxRepository interface {
	Update(ctx context.Context, event *Message) error
	GetNotPublishedMessages() ([]*Message, error)
}

var _ IOutboxRepository = &Repository{}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (*Repository, error) {
	if db == nil {
		return nil, errors.New("db")
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Update(ctx context.Context, outboxEvent *Message) error {
	err := r.db.WithContext(ctx).Save(&outboxEvent).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetNotPublishedMessages() ([]*Message, error) {
	var events []*Message
	result := r.db.
		Order("occurred_on_utc ASC").
		Limit(20).
		Where("processed_on_utc IS NULL").Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return events, nil
}
