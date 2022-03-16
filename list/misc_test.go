package list_test

import (
	"testing"

	"github.com/go-asset/generics/list"
	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	result := list.Head([]int{})

	assert.True(t, result.IsNothing())

	result = list.Head([]int{1, 2, 3})

	assert.False(t, result.IsNothing())
	assert.Equal(t, 1, result.Value())
}
