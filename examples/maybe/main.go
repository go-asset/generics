package main

import (
	"constraints"
	"fmt"

	"github.com/yitsushi/go1-18-experiments/pkg/data"
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
