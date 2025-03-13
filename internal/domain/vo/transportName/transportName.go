package transportName

import (
	"errors"
)

const (
	MIN = 3
	MAX = 100
)

type TransportName struct {
	value string
}

func New(name string) (TransportName, error) {
	if len(name) < MIN || len(name) > MAX {
		return TransportName{}, errors.New("TransportName is out of range")
	}

	return TransportName{name}, nil
}

func (t TransportName) Value() string { return t.value }

func (t TransportName) Equals(other TransportName) bool { return t == other }

func Fake() (TransportName, error) {
	return New("fake transport name") // TODO. random
}
