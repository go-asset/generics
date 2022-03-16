package list

import "github.com/go-asset/generics/maybe"

// Head extract the first element of a list as a Maybe, if the list is empty
// it's Nothing.
//
// It's categorized as "misc" because it's very limited when it can be useful,
// one example is chaining function calls with Maybe.
//
//     doSomething(FromMaybe(Head(myList), 0))
//
// But most of the time, it feels useless. Why implementing it anyway? Because
// why not. ;)
func Head[T any](list []T) maybe.Maybe[T] {
	if len(list) == 0 {
		return maybe.Nothing[T]()
	}

	return maybe.Just(list[0])
}
