package main

import "github.com/yitsushi/go1-18-experiments/pkg/data"

func readFromUser() chan data.Maybe[int] {
	ch := make(chan data.Maybe[int])

	myList := []data.Maybe[int]{
		data.Just(12),
		data.Just(2),
		data.Nothing[int](),
		data.Just(8),
	}

	go func() {
		for _, v := range myList {
			ch <- v
		}

		close(ch)
	}()

	return ch
}

func readFromUserStr() chan data.Maybe[string] {
	ch := make(chan data.Maybe[string])

	myList := []data.Maybe[string]{
		data.Just("asd"),
		data.Just("123"),
		data.Nothing[string](),
		data.Just(",./"),
	}

	go func() {
		for _, v := range myList {
			ch <- v
		}

		close(ch)
	}()

	return ch
}
