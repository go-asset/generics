package list

// Filter will return elements which return true based on the provided filtering function.
func Filter[T any](list []T, fn func(item T) bool) []T {
	var result []T
	for _, item := range list {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

// Some returns true if one or more items in the list satisfy the given function.
func Some[T any](list []T, fn func(item T) bool) bool {
	for _, item := range list {
		if fn(item) {
			return true
		}
	}
	return false
}

// All returns true if all the items in the list satisfy the given function.
func All[T any](list []T, fn func(item T) bool) bool {
	for _, item := range list {
		if !fn(item) {
			return false
		}
	}
	return true
}

// None returns true if no items in the list satisfy the given function.
func None[T any](list []T, fn func(item T) bool) bool {
	for _, item := range list {
		if fn(item) {
			return false
		}
	}
	return true
}
