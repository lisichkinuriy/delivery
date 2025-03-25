package queries

import (
	"errors"
	"gorm.io/gorm"
	"lisichkinuriy/delivery/internal/domain/order"
)

type IGetNotCompletedOrdersQueryHandler interface {
	Handle(GetNotCompletedOrdersQuery) (GetNotCompletedOrdersResponse, error)
}

type GetNotCompletedOrdersQueryHandler struct {
	db *gorm.DB
}

func NewGetNotCompletedOrdersQueryHandler(db *gorm.DB) (*GetNotCompletedOrdersQueryHandler, error) {
	if db == nil {
		return &GetNotCompletedOrdersQueryHandler{}, errors.New("db cannot be nil")
	}
	return &GetNotCompletedOrdersQueryHandler{db: db}, nil
}

func (q *GetNotCompletedOrdersQueryHandler) Handle(query GetNotCompletedOrdersQuery) (GetNotCompletedOrdersResponse, error) {
	if query.IsEmpty() {
		return GetNotCompletedOrdersResponse{}, errors.New("query is empty")
	}

	var orders []OrderResponse
	result := q.db.Raw("SELECT id, courier_id, location_x, location_y, status FROM orders where status!=?",
		order.StatusCompleted).Scan(&orders)

	if result.Error != nil {
		return GetNotCompletedOrdersResponse{}, result.Error
	}

	return GetNotCompletedOrdersResponse{Orders: orders}, nil
}
