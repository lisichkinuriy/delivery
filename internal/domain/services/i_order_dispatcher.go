package services

import (
	"lisichkinuriy/delivery/internal/domain/courier"
	"lisichkinuriy/delivery/internal/domain/order"
)

type IOrderDispatcher interface {
	Dispatch(order *order.Order, couriers []*courier.Courier) (*courier.Courier, error)
}
