package maybe_test

import (
	"testing"

	"github.com/go-asset/generics/maybe"
	"github.com/stretchr/testify/assert"
)

func TestMaybe(t *testing.T) {
	value := maybe.Just(12)

	assert.False(t, value.IsNothing())
	assert.Equal(t, 12, value.Value())

	value = maybe.Nothing[int]()

	assert.True(t, value.IsNothing())
}
