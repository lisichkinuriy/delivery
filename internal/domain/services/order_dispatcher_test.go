package services

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"lisichkinuriy/delivery/internal/domain/courier"
	"lisichkinuriy/delivery/internal/domain/order"
	"lisichkinuriy/delivery/internal/domain/vo"
	"testing"
)

func Test_CanDispatchOrder(t *testing.T) {

	speed, _ := vo.NewSpeed(1)
	name, _ := vo.FakeCourierName()
	transportName, _ := vo.FakeTransportName()

	location1, _ := vo.NewLocation(1, 1)
	c1, _ := courier.NewCourier(name, transportName, speed, location1)

	location2, _ := vo.NewLocation(2, 2)
	c2, _ := courier.NewCourier(name, transportName, speed, location2)

	couriers := []*courier.Courier{c1, c2}

	orderLocation, _ := vo.NewLocation(3, 3)
	o, _ := order.NewOrder(uuid.New(), orderLocation)

	dispatcher := NewOrderDispatcher()

	winner, err := dispatcher.Dispatch(o, couriers)
	assert.NoError(t, err)
	assert.Equal(t, c2, winner)
}
