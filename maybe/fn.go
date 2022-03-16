package maybe

// FromMaybe return the contents of the Maybe or a default value if it's
// Nothing.
func FromMaybe[T any](maybe Maybe[T], value T) T {
	if maybe.IsNothing() {
		return value
	}

	return maybe.Value()
}

// Values extracts values from Just items of a list while discarding all Nothing
// items.
func Values[T any](list []Maybe[T]) []T {
	result := []T{}

	for _, item := range list {
		if !item.IsNothing() {
			result = append(result, item.Value())
		}
	}

	return result
}
