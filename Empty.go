package collections

func Empty[T any]() Enumerable[T] {
	return emptyEnumerable[T]{}

}

type emptyEnumerable[T any] struct{}

func (emptyEnumerable[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T, 0)
	close(ch)
	return ch
}
