package courier

import (
	"errors"
	"github.com/google/uuid"
	"lisichkinuriy/delivery/internal/domain/courier/transport"
	"lisichkinuriy/delivery/internal/domain/vo"
)

type Courier struct {
	id        uuid.UUID
	name      vo.CourierName
	transport *transport.Transport
	location  vo.Location
	status    Status
}

func NewCourier(name vo.CourierName,
	transportName vo.TransportName,
	speed vo.Speed,
	location vo.Location) (*Courier, error) {

	if name.IsEmpty() {
		return nil, errors.New("name is empty")
	}
	if transportName.IsEmpty() {
		return nil, errors.New("transportName is empty")
	}
	if speed.IsEmpty() {
		return nil, errors.New("speed is empty")
	}
	if location.IsEmpty() {
		return nil, errors.New("location is empty")
	}

	t, err := transport.New(transportName, speed)
	if err != nil {
		return nil, err
	}

	return &Courier{
		id:        uuid.New(),
		name:      name,
		transport: t,
		location:  location,
		status:    StatusFree,
	}, nil
}

func (c *Courier) ID() uuid.UUID                   { return c.id }
func (c *Courier) Name() vo.CourierName            { return c.name }
func (c *Courier) Transport() *transport.Transport { return c.transport }
func (c *Courier) Location() vo.Location           { return c.location }

func (c *Courier) SetBusy() error {
	c.status = StatusBusy
	return nil
}

func (c *Courier) SetFree() error {
	c.status = StatusFree
	return nil
}

func (c *Courier) IsBusy() bool { return c.status == StatusBusy }
func (c *Courier) IsFree() bool { return c.status == StatusFree }

func (c *Courier) Equals(other *Courier) bool { return c.id == other.id }

func (c *Courier) CalculateTimeToLocation(target vo.Location) (float64, error) {
	if target.IsEmpty() {
		return 0, errors.New("target is empty")
	}

	distance := c.location.DistanceTo(target)

	time := float64(distance) / float64(c.transport.Speed().Value())
	return time, nil
}

func (c *Courier) Move(target vo.Location) error {
	if target.IsEmpty() {
		return errors.New("target is empty")
	}

	location, err := c.transport.Move(c.location, target)
	if err != nil {
		return err
	}
	c.location = location

	return nil
}

func (c *Courier) Status() Status {
	return c.status
}

func RestoreCourier(id uuid.UUID, name string,
	x int, y int, status Status,
	transportId uuid.UUID, transportName string, speed int) *Courier {
	t := transport.RestoreTransport(transportId, transportName, speed)
	cName, _ := vo.NewCourierName(name)
	l, _ := vo.NewLocation(x, y)

	return &Courier{
		id:        id,
		name:      cName,
		transport: t,
		location:  l,
		status:    status,
	}
}
