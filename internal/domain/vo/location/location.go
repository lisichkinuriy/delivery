package location

import (
	"errors"
	"lisichkinuriy/delivery/internal/utils"
	"math/rand"
	"time"
)

const (
	minY int = 1
	maxY int = 10
	minX int = 1
	maxX int = 10
)

type Location struct {
	x int
	y int
}

func New(x int, y int) (Location, error) {
	if x < minX || x > maxX {
		return Location{}, errors.New("x is out of range")
	}

	if y < minY || y > maxY {
		return Location{}, errors.New("y is out of range")
	}

	return Location{x, y}, nil
}

func Distance(l1, l2 Location) int {
	return utils.AbsInt(l1.X()-l2.X()) + utils.AbsInt(l1.Y()-l2.Y())
}

func (l Location) X() int                     { return l.x }
func (l Location) Y() int                     { return l.y }
func (l Location) Equals(other Location) bool { return l == other }

func Fake() (Location, error) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	x := r.Intn(maxX) + 1
	y := r.Intn(maxY) + 1

	return New(x, y)
}
