package either

// Left creates a Left value.
func Left[L any, R any](value L) Either[L, R] {
	return Either[L, R]{left: value}
}

// Right creates a Right value.
func Right[L any, R any](value R) Either[L, R] {
	return Either[L, R]{right: value, isRight: true}
}
