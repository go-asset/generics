package list_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-asset/generics/list"
)

func TestFoldIter(t *testing.T) {
	ch := make(chan int)

	myList := []int{12, 23, 8, 14, 44, 1, 3, 52}

	go func() {
		for _, v := range myList {
			ch <- v
		}

		close(ch)
	}()

	result := list.FoldlIter(0, ch, func(c int, value int) int {
		return c + value
	})

	assert.Equal(t, 157, result)
}

func TestFoldl(t *testing.T) {
	myList := []byte("asdfghjkl")

	result := list.Foldl("", myList, func(c string, value byte) string {
		return c + string(value)
	})

	assert.Equal(t, "asdfghjkl", result)
}

func TestFoldr(t *testing.T) {
	myList := []byte("asdfghjkl")

	result := list.Foldr("", myList, func(c string, value byte) string {
		return c + string(value)
	})

	assert.Equal(t, "lkjhgfdsa", result)
}
