package vo

import (
	"errors"
)

const (
	MIN_CN_LEN = 3
	MAX_CN_LEN = 100
)

type CourierName struct {
	value string
}

func NewCourierName(name string) (CourierName, error) {
	if len(name) < MIN_CN_LEN || len(name) > MAX_CN_LEN {
		return CourierName{}, errors.New("CourierName is out of range")
	}

	return CourierName{name}, nil
}

func (t CourierName) Value() string { return t.value }

func (t CourierName) Equals(other CourierName) bool { return t == other }

func (t CourierName) IsEmpty() bool {
	return t.Equals(CourierName{})
}

func FakeCourierName() (CourierName, error) {
	return NewCourierName("Vasya")
}
