package maybe

// Maybe is a data type that can hold a value or be Nothing.
type Maybe[T any] struct {
	value     T
	isNothing bool
}

// Value returns with the value of the type.
func (m Maybe[T]) Value() T {
	return m.value
}

// IsNothing returns true if it's a Nothing.
func (m Maybe[T]) IsNothing() bool {
	return m.isNothing
}
