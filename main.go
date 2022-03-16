package main

import (
	"log"
	"time"

	"github.com/go-asset/generics/list"
)

func main() {
	result, cancel := list.MapStream(
		[]int{1, 2, 3, 4},
		func(v int) int {
			time.Sleep(1 * time.Second)
			return v * 2
		},
	)
	defer cancel()

	endResult := []int{}

	for v := range result {
		log.Println(v)

		if v == 4 {
			cancel()
		}

		endResult = append(endResult, v)
	}
}
