package services

import (
	"errors"
	"lisichkinuriy/delivery/internal/domain/courier"
	"lisichkinuriy/delivery/internal/domain/order"
	"math"
)

type OrderDispatcher struct{}

func NewOrderDispatcher() *OrderDispatcher { return &OrderDispatcher{} }

func (p *OrderDispatcher) Dispatch(order *order.Order, couriers []*courier.Courier) (*courier.Courier, error) {
	if order == nil {
		return nil, errors.New("nil order")
	}

	if couriers == nil || len(couriers) == 0 {
		return nil, errors.New("empty couriers")
	}

	var bestCourier *courier.Courier = nil
	minTime := math.MaxFloat64

	for _, c := range couriers {
		courierTime, err := c.CalculateTimeToLocation(order.Location())
		if err != nil {
			return nil, err
		}

		if courierTime < minTime {
			minTime = courierTime
			bestCourier = c
		}
	}

	if bestCourier == nil {
		return nil, errors.New("cannot find the best courier")
	}
	err := order.Assign(bestCourier)
	if err != nil {
		return nil, err
	}
	err = bestCourier.SetBusy()
	if err != nil {
		return nil, err
	}

	return bestCourier, nil
}
