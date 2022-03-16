package list

// Map applies fn function to all elements of the list and returns with the list
// of the results.
func Map[T any, R any](list []T, fn func(T) R) []R {
	result := []R{}

	for _, item := range list {
		result = append(result, fn(item))
	}

	return result
}

// MapIter applies fn function to all elements of a chan and returns with the
// list of the results.
func MapIter[T any, R any](list <-chan T, fn func(T) R) []R {
	result := []R{}

	for item := range list {
		result = append(result, fn(item))
	}

	return result
}

// MapStream applies fn function to all elements of a list and streams back the
// results in the result chan. Can be cancelled with the cancel function.
func MapStream[T any, R any](list []T, fn func(T) R) (<-chan R, func()) {
	result := make(chan R)
	cancel := make(chan bool)
	cancelFn := func() { close(cancel) }

	go func() {
		defer close(result)

		for _, item := range list {
			select {
			case <-cancel:
				return
			default:
			}

			result <- fn(item)
		}

	}()

	return result, cancelFn
}

// MapIterStream applies fn function to all elements of a chan and streams back
// the results in the result chan. Can be cancelled with the cancel function.
func MapIterStream[T any, R any](list <-chan T, fn func(T) R) (<-chan R, func()) {
	result := make(chan R)
	cancel := make(chan bool)
	cancelFn := func() { close(cancel) }

	go func() {
		defer close(result)

		for item := range list {
			select {
			case <-cancel:
				return
			default:
			}

			result <- fn(item)
		}

	}()

	return result, cancelFn
}
