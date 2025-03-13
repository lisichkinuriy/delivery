package transport

import (
	"github.com/stretchr/testify/assert"
	"lisichkinuriy/delivery/internal/domain/vo/location"
	"lisichkinuriy/delivery/internal/domain/vo/speed"
	"lisichkinuriy/delivery/internal/domain/vo/transportName"
	"testing"
)

func TestNew(t *testing.T) {
	name, _ := transportName.Fake()
	s, _ := speed.Fake()
	transport, err := New(name, s)

	assert.NoError(t, err)
	assert.Equal(t, name, transport.Name())
	assert.Equal(t, s, transport.Speed())
}

func TestTransport_Move(t *testing.T) {
	name, _ := transportName.New("Pedestrian")
	s, _ := speed.New(1)

	pedestrian, err := New(name, s)
	assert.NoError(t, err)

	name, _ = transportName.New("Bicycle")
	s, _ = speed.New(2)
	bicycle, err := New(name, s)
	assert.NoError(t, err)

	location11, _ := location.New(1, 1)
	location12, _ := location.New(1, 2)
	location13, _ := location.New(1, 3)
	location15, _ := location.New(1, 5)
	location22, _ := location.New(2, 2)
	location32, _ := location.New(3, 2)
	location42, _ := location.New(4, 2)

	tests := []struct {
		name      string
		transport *Transport
		start     location.Location
		target    location.Location
		expected  location.Location
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
