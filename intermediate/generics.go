package main

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	items []T
}

func main() {

}
