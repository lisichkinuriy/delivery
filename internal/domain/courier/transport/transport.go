package transport

import (
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/domain/vo"
	"math"
)

type Transport struct {
	id    uuid.UUID
	name  vo.TransportName
	speed vo.Speed
}

func New(name vo.TransportName, speed vo.Speed) (*Transport, error) {

	return &Transport{
		id:    uuid.New(),
		name:  name,
		speed: speed,
	}, nil
}

func (t *Transport) ID() uuid.UUID          { return t.id }
func (t *Transport) Name() vo.TransportName { return t.name }
func (t *Transport) Speed() vo.Speed        { return t.speed }

func (t *Transport) Move(current vo.Location, target vo.Location) (vo.Location, error) {
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

	newLocation, err := vo.NewLocation(newX, newY)
	if err != nil {
		return current, err
	}
	return newLocation, nil
}
