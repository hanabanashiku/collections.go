package collections

func Yield[T any](enumerator chan<- *T, values Enumerator[T]) bool {
	current, ok := <-values

	if ok {
		enumerator <- current
		return true
	}

	return false
}

func YieldRange[T any](enumerator chan<- *T, enumerable Enumerable[T]) {
	for current := range enumerable.GetEnumerator() {
		enumerator <- current
	}
}
