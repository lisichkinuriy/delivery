package courier

import (
	"github.com/stretchr/testify/assert"
	"lisichkinuriy/delivery/internal/domain/vo"
	"testing"
)

func Test_CourierShouldBeCorrectWhenParamsAreCorrectOnCreated(t *testing.T) {
	// Arrange

	transport_name, _ := vo.FakeTransportName()
	speed, _ := vo.FakeSpeed()
	name := "Велосипедист"
	l := vo.MinLocation()

	// Act
	courier, err := NewCourier(name, transport_name, speed, l)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, courier)
	assert.Equal(t, name, courier.Name())
	assert.Equal(t, transport_name, courier.Transport().Name())
	assert.Equal(t, speed, courier.Transport().Speed())
	assert.Equal(t, l, courier.Location())
}

func Test_CourierCanCalculateTimeToLocation(t *testing.T) {
	// Изначальная точка курьера: [1,1]
	// Целевая точка: [5,10]
	// Количество шагов, необходимое курьеру: 13 (4 по горизонтали и 9 по вертикали)
	// Скорость транспорта (велосипедиста): 2 шага в 1 такт
	// Время подлета: 13/2 = 6.5 тактов потребуется курьеру, чтобы доставить заказ

	transport_name, _ := vo.FakeTransportName()
	speed, _ := vo.NewSpeed(2)
	name := "Велосипедист"
	l := vo.MinLocation()

	// Arrange
	courier, err := NewCourier(name, transport_name, speed, l)
	assert.NoError(t, err)
	target, err := vo.NewLocation(5, 10)
	assert.NoError(t, err)

	// Act
	time, err := courier.CalculateMovesToLocation(target)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 6.5, time)
}
