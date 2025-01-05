package collections

func Where[T any](enumerable Enumerable[T], predicate func(*T) bool) Enumerable[T] {
	return whereEnumerable[T]{
		values:    enumerable,
		predicate: predicate,
	}
}

type whereEnumerable[T any] struct {
	values    Enumerable[T]
	predicate func(*T) bool
}
type whereEnumerator[T any] struct {
	enumerator Enumerator[T]
	predicate  func(*T) bool
}

func (enumerable whereEnumerable[T]) GetEnumerator() Enumerator[T] {
	return whereEnumerator[T]{
		enumerator: enumerable.values.GetEnumerator(),
		predicate:  enumerable.predicate,
	}
}

func (enumerator whereEnumerator[T]) MoveNext() bool {
	for enumerator.enumerator.MoveNext() {
		if enumerator.predicate(enumerator.enumerator.Current()) {
			return true
		}
	}

	return false
}

func (enumerator whereEnumerator[T]) Current() *T {
	return enumerator.enumerator.Current()
}

func (enumerator whereEnumerator[T]) Reset() {
	enumerator.enumerator.Reset()
}
