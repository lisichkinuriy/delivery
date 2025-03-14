package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	value := MIN_SPEED

	speed, err := NewSpeed(value)
	assert.NoError(t, err)
	assert.Equal(t, value, speed.Value())
}

func TestSpeed_Fake(t *testing.T) {
	_, err := FakeSpeed()
	assert.NoError(t, err)
}

func TestConstraints(t *testing.T) {
	value := MIN_SPEED - 1
	_, err := NewSpeed(value)
	assert.NotNil(t, err)
}
