package transport

import (
	"github.com/stretchr/testify/assert"
	"lisichkinuriy/delivery/internal/domain/vo"
	"testing"
)

func TestNew(t *testing.T) {
	name, _ := vo.FakeTransportName()
	s, _ := vo.FakeSpeed()
	transport, err := New(name, s)

	assert.NoError(t, err)
	assert.Equal(t, name, transport.Name())
	assert.Equal(t, s, transport.Speed())
}

func TestTransport_Move(t *testing.T) {
	name, _ := vo.NewTransportName("Pedestrian")
	s, _ := vo.NewSpeed(1)

	pedestrian, err := New(name, s)
	assert.NoError(t, err)

	name, _ = vo.NewTransportName("Bicycle")
	s, _ = vo.NewSpeed(2)
	bicycle, err := New(name, s)
	assert.NoError(t, err)

	location11, _ := vo.NewLocation(1, 1)
	location12, _ := vo.NewLocation(1, 2)
	location13, _ := vo.NewLocation(1, 3)
	location15, _ := vo.NewLocation(1, 5)
	location22, _ := vo.NewLocation(2, 2)
	location32, _ := vo.NewLocation(3, 2)
	location42, _ := vo.NewLocation(4, 2)

	tests := []struct {
		name      string
		transport *Transport
		start     vo.Location
		target    vo.Location
		expected  vo.Location
	}{
		// Пешеход
		{"Pedestrian Same Location", pedestrian, location11, location11, location11},
		{"Pedestrian Move Up", pedestrian, location11, location12, location12},
		{"Pedestrian Move Up Limited", pedestrian, location11, location15, location12},
		{"Pedestrian Move Right", pedestrian, location22, location32, location32},
		{"Pedestrian Move Down", pedestrian, location12, location11, location11},
		{"Pedestrian Move Left", pedestrian, location22, location12, location12},

		// Велосипедист
		{"Bicycle Same Location", bicycle, location11, location11, location11},
		{"Bicycle Move Up", bicycle, location11, location13, location13},
		{"Bicycle Move Up Limited", bicycle, location11, location15, location13},
		{"Bicycle Move Right", bicycle, location22, location42, location42},
		{"Bicycle Move Down", bicycle, location13, location11, location11},
		{"Bicycle Move Left", bicycle, location32, location12, location12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newLocation, err := tt.transport.Move(tt.start, tt.target)
			assert.NoError(t, err)

			if newLocation != tt.expected {
				t.Errorf("got %v, want %v", newLocation, tt.expected)
			}
		})
	}
}
