package maybe

type Maybe[T any] struct {
	value     T
	isNothing bool
}

func (m Maybe[T]) Value() T {
	return m.value
}

func (m Maybe[T]) IsNothing() bool {
	return m.isNothing
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{isNothing: true}
}

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{value: value}
}
