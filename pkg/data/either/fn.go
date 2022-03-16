package either

// FromLeft return the contents of a Left-value or a default value otherwise.
func FromLeft[L any, R any](either Either[L, R], value L) L {
	if either.IsRight() {
		return value
	}

	return either.Left()
}

// FromRight return the contents of a Right-value or a default value otherwise.
func FromRight[L any, R any](either Either[L, R], value R) R {
	if either.IsRight() {
		return either.Right()
	}

	return value
}

// Lefts extracts from a list of Either all the Left elements.
// All the Left elements are extracted in order.
func Lefts[L any, R any](eithers []Either[L, R]) []L {
	lefts := []L{}

	for _, either := range eithers {
		if either.IsLeft() {
			lefts = append(lefts, either.Left())
		}
	}

	return lefts
}

// Rights extracts from a list of Either all the Right elements.
// All the Right elements are extracted in order.
func Rights[L any, R any](eithers []Either[L, R]) []R {
	rights := []R{}

	for _, either := range eithers {
		if either.IsRight() {
			rights = append(rights, either.Right())
		}
	}

	return rights
}
