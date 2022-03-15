package main

import (
	"fmt"
	"log"

	"golang.org/x/exp/constraints"

	"github.com/yitsushi/go-generics/pkg/data"
)

func add[L any, R constraints.Ordered](c data.Maybe[R], value data.Either[L, R]) data.Maybe[R] {
	if !value.IsRight() {
		log.Printf("Skip: %v\n", value.Left())

		return c
	}

	if c.IsNothing() {
		return data.Just(value.Right())
	}

	return data.Just(c.Value() + value.Right())
}

func main() {
	result := data.FoldlIter(
		data.Nothing[int](),
		readFromUser(),
		add[string, int],
	)

	fmt.Println(result)
}
