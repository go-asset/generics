package list_test

import (
	"testing"

	"github.com/go-asset/generics/list"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	result := list.Map([]int{1, 2, 3}, func(v int) int { return v * 2 })

	assert.Equal(t, []int{2, 4, 6}, result)
}

func TestMapIter(t *testing.T) {
	ch := make(chan int)

	go func() {
		for _, v := range []int{1, 2, 3} {
			ch <- v
		}

		close(ch)
	}()

	result := list.MapIter(ch, func(v int) int { return v * 2 })

	assert.Equal(t, []int{2, 4, 6}, result)
}

func TestMapStream(t *testing.T) {
	result, cancel := list.MapStream([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})
	defer cancel()

	endResult := []int{}

	for v := range result {
		endResult = append(endResult, v)
	}

	assert.Equal(t, []int{2, 4, 6}, endResult)
}

func TestMapIterStream(t *testing.T) {
	ch := make(chan int)

	go func() {
		for _, v := range []int{1, 2, 3} {
			ch <- v
		}

		close(ch)
	}()

	result, cancel := list.MapIterStream(ch, func(v int) int {
		return v * 2
	})
	defer cancel()

	endResult := []int{}

	for v := range result {
		endResult = append(endResult, v)
	}

	assert.Equal(t, []int{2, 4, 6}, endResult)
}
