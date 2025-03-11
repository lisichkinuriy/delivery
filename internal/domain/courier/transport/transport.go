package transport

import (
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/domain/vo/location"
	"lisichkinuriy/delivery/internal/domain/vo/speed"
	"lisichkinuriy/delivery/internal/domain/vo/transportName"
	"math"
)

type Transport struct {
	id    uuid.UUID
	name  transportName.TransportName
	speed speed.Speed
}

func New(name transportName.TransportName, speed speed.Speed) (*Transport, error) {

	return &Transport{
		id:    uuid.New(),
		name:  name,
		speed: speed,
	}, nil
}

func (t *Transport) ID() uuid.UUID                     { return t.id }
func (t *Transport) Name() transportName.TransportName { return t.name }
func (t *Transport) Speed() speed.Speed                { return t.speed }

func (t *Transport) Move(current location.Location, target location.Location) (location.Location, error) {
	if current == target {
		return current, nil // Уже на месте
	}

	dx := float64(target.X() - current.X())
	dy := float64(target.Y() - current.Y())
	remainingRange := float64(t.speed.Value())

	if math.Abs(dx) > remainingRange {
		dx = math.Copysign(remainingRange, dx)
	}
	remainingRange -= math.Abs(dx)

	if math.Abs(dy) > remainingRange {
		dy = math.Copysign(remainingRange, dy)
	}

	newX := current.X() + int(dx)
	newY := current.Y() + int(dy)

	newLocation, err := location.New(newX, newY)
	if err != nil {
		return current, err
	}
	return newLocation, nil
}
