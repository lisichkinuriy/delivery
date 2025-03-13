package transportName

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	value := "hello"

	name, err := New(value)
	assert.NoError(t, err)
	assert.Equal(t, value, name.Value())
}

func TestSpeed_Fake(t *testing.T) {
	_, err := Fake()
	assert.NoError(t, err)
}

func TestConstraints(t *testing.T) {
	value := ""
	_, err := New(value)
	assert.NotNil(t, err)
}
