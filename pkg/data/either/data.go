package either

// Either type represents values with two possibilities:
// a value of type Either a b is either Left a or Right b.
type Either[L any, R any] struct {
	left  L
	right R

	isRight bool
}

// Left returns with the Left value.
func (e Either[L, R]) Left() L {
	return e.left
}

// Right returns with the Right value.
func (e Either[L, R]) Right() R {
	return e.right
}

// IsRight returns true if it's a Right value.
func (e Either[L, R]) IsRight() bool {
	return e.isRight
}

// IsLeft returns true if it's a Left value.
func (e Either[L, R]) IsLeft() bool {
	return !e.isRight
}
