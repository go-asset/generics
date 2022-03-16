package main

import (
	"fmt"

	"golang.org/x/exp/constraints"

	"github.com/go-asset/generics/pkg/data"
)

func add[T constraints.Ordered](c data.Maybe[T], value data.Maybe[T]) data.Maybe[T] {
	if c.IsNothing() {
		return value
	}

	if value.IsNothing() {
		return c
	}

	return data.Just(c.Value() + value.Value())
}

func main() {
	result := data.FoldlIter(
		data.Just(0),
		readFromUser(),
		add[int],
	)

	resultStr := data.FoldlIter(
		data.Just(""),
		readFromUserStr(),
		add[string],
	)

	fmt.Println(result)
	fmt.Println(resultStr)
}
