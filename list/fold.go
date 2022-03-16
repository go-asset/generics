package list

// FoldlIter is a left-associative fold of a chan.
func FoldlIter[T any, R any](init R, list <-chan T, fold func(carry R, next T) R) R {
	for value := range list {
		init = fold(init, value)
	}

	return init
}

// Foldl is a left-associative fold of a list.
func Foldl[T any, R any](init R, list []T, fold func(carry R, next T) R) R {
	for _, value := range list {
		init = fold(init, value)
	}

	return init
}

// Foldr is a right-associative fold of a list.
func Foldr[T any, R any](init R, list []T, fold func(carry R, next T) R) R {
	for idx := len(list) - 1; idx >= 0; idx-- {
		init = fold(init, list[idx])
	}

	return init
}
