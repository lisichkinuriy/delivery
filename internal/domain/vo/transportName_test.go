package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewTransportName(t *testing.T) {
	value := "hello"

	name, err := NewTransportName(value)
	assert.NoError(t, err)
	assert.Equal(t, value, name.Value())
}

func Test_TransportNameFake(t *testing.T) {
	_, err := FakeSpeed()
	assert.NoError(t, err)
}

func Test_TransportNameConstraints(t *testing.T) {
	value := ""
	_, err := NewTransportName(value)
	assert.NotNil(t, err)
}
