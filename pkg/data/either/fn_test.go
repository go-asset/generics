package either_test

import (
	"testing"

	"github.com/go-asset/generics/pkg/data/either"
	"github.com/stretchr/testify/assert"
)

func TestFromLeft(t *testing.T) {
	value := either.Left[string, int]("left value")

	assert.Equal(t, "left value", either.FromLeft(value, "default"))

	value = either.Right[string](12)

	assert.Equal(t, "default", either.FromLeft(value, "default"))
}

func TestFromRight(t *testing.T) {
	value := either.Left[string, int]("left value")

	assert.Equal(t, 99, either.FromRight(value, 99))

	value = either.Right[string](12)

	assert.Equal(t, 12, either.FromRight(value, 99))
}

func TestLefts(t *testing.T) {
	values := []either.Either[string, int]{
		either.Left[string, int]("a"),
		either.Left[string, int]("b"),
		either.Right[string](1),
		either.Left[string, int]("c"),
		either.Right[string](2),
		either.Right[string](3),
	}

	assert.Equal(t, []string{"a", "b", "c"}, either.Lefts(values))
}

func TestLefts_allLeft(t *testing.T) {
	values := []either.Either[string, int]{
		either.Left[string, int]("a"),
		either.Left[string, int]("b"),
		either.Left[string, int]("c"),
	}

	assert.Equal(t, []string{"a", "b", "c"}, either.Lefts(values))
}

func TestLefts_allRights(t *testing.T) {
	values := []either.Either[string, int]{
		either.Right[string](1),
		either.Right[string](2),
		either.Right[string](3),
	}

	assert.Equal(t, []string{}, either.Lefts(values))
}

func TestLefts_empty(t *testing.T) {
	values := []either.Either[string, int]{}

	assert.Equal(t, []string{}, either.Lefts(values))
}

func TestRights(t *testing.T) {
	values := []either.Either[string, int]{
		either.Left[string, int]("a"),
		either.Left[string, int]("b"),
		either.Right[string](1),
		either.Left[string, int]("c"),
		either.Right[string](2),
		either.Right[string](3),
	}

	assert.Equal(t, []int{1, 2, 3}, either.Rights(values))
}

func TestRights_allLeft(t *testing.T) {
	values := []either.Either[string, int]{
		either.Left[string, int]("a"),
		either.Left[string, int]("b"),
		either.Left[string, int]("c"),
	}

	assert.Equal(t, []int{}, either.Rights(values))
}

func TestRights_allRights(t *testing.T) {
	values := []either.Either[string, int]{
		either.Right[string](1),
		either.Right[string](2),
		either.Right[string](3),
	}

	assert.Equal(t, []int{1, 2, 3}, either.Rights(values))
}

func TestRights_empty(t *testing.T) {
	values := []either.Either[string, int]{}

	assert.Equal(t, []int{}, either.Rights(values))
}
