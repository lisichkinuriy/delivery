package vo

import (
	"errors"
)

const (
	MIN_TN_LEN = 3
	MAX_TN_LEN = 100
)

type TransportName struct {
	value string
}

func NewTransportName(name string) (TransportName, error) {
	if len(name) < MIN_TN_LEN || len(name) > MAX_TN_LEN {
		return TransportName{}, errors.New("TransportName is out of range")
	}

	return TransportName{name}, nil
}

func (t TransportName) Value() string { return t.value }

func (t TransportName) Equals(other TransportName) bool { return t == other }

func (t TransportName) IsEmpty() bool {
	return t.Equals(TransportName{})
}

func FakeTransportName() (TransportName, error) {
	return NewTransportName("fake transport name")
}
