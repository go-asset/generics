package maybe_test

import (
	"testing"

	"github.com/go-asset/generics/maybe"
	"github.com/stretchr/testify/assert"
)

func TestFromMaybe(t *testing.T) {
	value := maybe.Just(12)

	assert.Equal(t, 12, maybe.FromMaybe(value, 99))

	value = maybe.Nothing[int]()

	assert.Equal(t, 99, maybe.FromMaybe(value, 99))
}

func TestValues(t *testing.T) {
	myList := []maybe.Maybe[int]{
		maybe.Just(1),
		maybe.Nothing[int](),
		maybe.Just(2),
		maybe.Just(3),
		maybe.Nothing[int](),
		maybe.Just(4),
	}

	assert.Equal(t, []int{1, 2, 3, 4}, maybe.Values(myList))
}

func TestValues_allJust(t *testing.T) {
	myList := []maybe.Maybe[int]{
		maybe.Just(1),
		maybe.Just(2),
		maybe.Just(3),
		maybe.Just(4),
	}

	assert.Equal(t, []int{1, 2, 3, 4}, maybe.Values(myList))
}

func TestValues_allNothing(t *testing.T) {
	myList := []maybe.Maybe[int]{
		maybe.Nothing[int](),
		maybe.Nothing[int](),
	}

	assert.Equal(t, []int{}, maybe.Values(myList))
}

func TestValues_empty(t *testing.T) {
	myList := []maybe.Maybe[int]{}

	assert.Equal(t, []int{}, maybe.Values(myList))
}
