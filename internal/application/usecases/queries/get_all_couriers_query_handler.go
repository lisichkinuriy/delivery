package queries

import (
	"errors"
	"gorm.io/gorm"
)

type IGetAllCouriersQueryHandler interface {
	Handle(GetAllCouriersQuery) (GetAllCouriersResponse, error)
}

type GetAllCouriersQueryHandler struct {
	db *gorm.DB
}

func NewGetAllCouriersQueryHandler(db *gorm.DB) (*GetAllCouriersQueryHandler, error) {
	if db == nil {
		return &GetAllCouriersQueryHandler{}, errors.New("db cannot be nil")
	}
	return &GetAllCouriersQueryHandler{db: db}, nil
}

func (q *GetAllCouriersQueryHandler) Handle(query GetAllCouriersQuery) (GetAllCouriersResponse, error) {
	if query.IsEmpty() {
		return GetAllCouriersResponse{}, errors.New("query is empty")
	}

	var couriers []CourierResponse
	result := q.db.Raw("SELECT id,name, location_x, location_y FROM couriers").Scan(&couriers)

	if result.Error != nil {
		return GetAllCouriersResponse{}, result.Error
	}

	return GetAllCouriersResponse{Couriers: couriers}, nil
}
