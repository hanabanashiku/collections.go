package collections

func Select[TValue, TNext any](enumerable Enumerable[TValue], mapper func(*TValue) TNext) Enumerable[TNext] {
	if enumerable == nil {
		return Empty[TNext]()
	}

	if mapper == nil {
		panic("mapper cannot be nil")
	}

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
	ch := make(chan *TNext)

	go func() {
		defer close(ch)
		for current := range enumerable.values.GetEnumerator() {
			next := enumerable.mapper(current)
			ch <- &next
		}

	}()

	return ch
}
