package either_test

import (
	"testing"

	"github.com/go-asset/generics/either"
	"github.com/stretchr/testify/assert"
)

func TestEither(t *testing.T) {
	value := either.Left[string, int]("left value")

	assert.True(t, value.IsLeft())
	assert.Equal(t, "left value", value.Left())

	value = either.Right[string](12)

	assert.True(t, value.IsRight())
	assert.Equal(t, 12, value.Right())
}
