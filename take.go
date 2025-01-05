package collections

func Take[T any](source Enumerable[T], count int) Enumerable[T] {
	return takeEnumerable[T]{
		source: source,
		count:  count,
	}
}

type takeEnumerable[T any] struct {
	source Enumerable[T]
	count  int
}

func (enumerable takeEnumerable[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T)

	go func() {
		i := 0
		for current := range enumerable.source.GetEnumerator() {
			if i >= enumerable.count {
				close(ch)
				return
			}

			ch <- current
			i++
		}

		close(ch)
	}()

	return ch
}
