package collections

func Empty[T any]() Enumerable[T] {
	return emptyEnumerable[T]{}

}

type emptyEnumerable[T any] struct{}
type emptyEnumerator[T any] struct{}

func (emptyEnumerable[T]) GetEnumerator() Enumerator[T] {
	return emptyEnumerable[T]{}
}

func (enumerator emptyEnumerable[T]) Current() *T {
	return nil
}

func (enumerator emptyEnumerable[T]) MoveNext() bool {
	return false
}

func (enumerator emptyEnumerable[T]) Reset() {
}
