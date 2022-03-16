package main

import "github.com/go-asset/generics/pkg/data"

func readFromUser() chan data.Either[string, int] {
	ch := make(chan data.Either[string, int])

	myList := []data.Either[string, int]{
		data.Right[string](12),
		data.Right[string](2),
		data.Left[string, int]("asd"),
		data.Right[string](8),
	}

	go func() {
		for _, v := range myList {
			ch <- v
		}

		close(ch)
	}()

	return ch
}
