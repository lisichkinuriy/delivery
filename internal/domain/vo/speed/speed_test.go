package speed

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	value := MIN

	speed, err := New(value)
	assert.NoError(t, err)
	assert.Equal(t, value, speed.Value())
}

func TestSpeed_Fake(t *testing.T) {
	_, err := Fake()
	assert.NoError(t, err)
}

func TestConstraints(t *testing.T) {
	value := MIN - 1
	_, err := New(value)
	assert.NotNil(t, err)
}
