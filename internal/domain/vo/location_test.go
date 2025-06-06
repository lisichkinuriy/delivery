package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LocationCanBeCreated(t *testing.T) {
	x := 1
	y := 2
	l, err := NewLocation(x, y)

	assert.NoError(t, err)
	assert.Equal(t, x, l.X())
	assert.Equal(t, y, l.Y())
}

func Test_CanCalcDistance(t *testing.T) {

	x1 := minX
	y1 := minY
	x2 := 5
	y2 := minY
	expected := 4

	// TODO: заделать цикл

	l1, _ := NewLocation(x1, y1)
	l2, _ := NewLocation(x2, y2)

	d := Distance(l1, l2)

	assert.Equal(t, expected, d)
}

func Test_FakeLocation(t *testing.T) {
	for i := 0; i < 100; i++ {
		_, err := FakeLocation()
		assert.NoError(t, err)
	}
}
