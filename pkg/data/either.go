package data

type Either[L any, R any] struct {
	left  L
	right R

	isRight bool
}

func (e Either[L, R]) Left() L {
	return e.left
}

func (e Either[L, R]) Right() R {
	return e.right
}

func (e Either[L, R]) IsRight() bool {
	return e.isRight
}

func Left[L any, R any](value L) Either[L, R] {
	return Either[L, R]{left: value}
}

func Right[L any, R any](value R) Either[L, R] {
	return Either[L, R]{right: value, isRight: true}
}
