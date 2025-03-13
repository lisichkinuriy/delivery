package speed

import "errors"

const (
	MIN = 1
	MAX = 3
)

type Speed struct {
	value int
}

func New(value int) (Speed, error) {
	if value < MIN || value > MAX {
		return Speed{}, errors.New("Speed must be between 1 and 3")
	}
	return Speed{value}, nil
}

func (s Speed) Value() int { return s.value }

func (s Speed) Equals(other Speed) bool { return s == other }

func Fake() (Speed, error) {
	return New(MAX) // TODO. random
}
