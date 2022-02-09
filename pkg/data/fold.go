package data

func FoldlIter[T any, R any](init R, list <-chan T, fold func(carry R, next T) R) R {
	for value := range list {
		init = fold(init, value)
	}

	return init
}

func Foldl[T any, R any](init R, list []T, fold func(carry R, next T) R) R {
	for _, value := range list {
		init = fold(init, value)
	}

	return init
}

func Foldr[T any, R any](init R, list []T, fold func(carry R, next T) R) R {
	for idx := len(list) - 1; idx >= 0; idx-- {
		init = fold(init, list[idx])
	}

	return init
}
