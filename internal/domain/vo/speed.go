package vo

import "errors"

const (
	MIN_SPEED = 1
	MAX_SPEED = 3
)

type Speed struct {
	value int
}

func NewSpeed(value int) (Speed, error) {
	if value < MIN_SPEED || value > MAX_SPEED {
		return Speed{}, errors.New("Speed must be between 1 and 3")
	}
	return Speed{value}, nil
}

func (s Speed) Value() int { return s.value }

func (s Speed) Equals(other Speed) bool { return s == other }

func (s Speed) IsEmpty() bool {
	return s.Equals(Speed{})
}

func FakeSpeed() (Speed, error) {
	return NewSpeed(MAX_SPEED)
}
