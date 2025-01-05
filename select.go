package collections

func Select[TValue, TNext any](enumerable Enumerable[TValue], mapper func(*TValue) TNext) Enumerable[TNext] {
	return selectEnumerable[TValue, TNext]{
		values: enumerable,
		mapper: mapper,
	}
}

type selectEnumerable[TValue, TNext any] struct {
	values Enumerable[TValue]
	mapper func(*TValue) TNext
}

func (enumerable selectEnumerable[TValue, TNext]) GetEnumerator() Enumerator[TNext] {
	return selectEnumerator[TValue, TNext]{
		enumerator: enumerable.values.GetEnumerator(),
		mapper:     enumerable.mapper,
	}
}

type selectEnumerator[TValue, TNext any] struct {
	enumerator Enumerator[TValue]
	mapper     func(*TValue) TNext
}

func (enumerator selectEnumerator[TValue, TNext]) MoveNext() bool {
	return enumerator.enumerator.MoveNext()
}

func (enumerator selectEnumerator[TValue, TNext]) Current() *TNext {
	current := enumerator.enumerator.Current()
	if current == nil {
		return nil
	}

	mapped := enumerator.mapper(current)
	return &mapped
}

func (enumerator selectEnumerator[TValue, TNext]) Reset() {
	enumerator.enumerator.Reset()
}
