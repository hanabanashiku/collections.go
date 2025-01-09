package collections

func Where[T any](enumerable Enumerable[T], predicate func(*T) bool) Enumerable[T] {
	if enumerable == nil {
		return Empty[T]()
	}

	if predicate == nil {
		panic("Predicate cannot be nil")
	}

	return whereEnumerable[T]{
		values:    enumerable,
		predicate: predicate,
	}
}

type whereEnumerable[T any] struct {
	values    Enumerable[T]
	predicate func(*T) bool
}

func (enumerable whereEnumerable[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T)

	go func() {
		defer close(ch)
		for current := range enumerable.GetEnumerator() {
			if enumerable.predicate(current) {
				ch <- current
			}
		}
	}()

	return ch
}
