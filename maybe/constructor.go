package maybe

// Nothing returns with Maybe that's Nothing.
func Nothing[T any]() Maybe[T] {
	return Maybe[T]{isNothing: true}
}

// Just returns with Maybe that has a value.
func Just[T any](value T) Maybe[T] {
	return Maybe[T]{value: value}
}
