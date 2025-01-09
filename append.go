package collections

func Append[T any](enumerable Enumerable[T], values Enumerable[T]) Enumerable[T] {
	return Insert(enumerable, values, -1)
}

func Prepend[T any](enumerable Enumerable[T], values Enumerable[T]) Enumerable[T] {
	return Insert(enumerable, values, 0)
}

func Insert[T any](enumerable Enumerable[T], values Enumerable[T], index int) Enumerable[T] {
	return insertAtEnumerable[T]{
		values:   enumerable,
		toInsert: values,
		index:    index,
	}
}

type insertAtEnumerable[T any] struct {
	values   Enumerable[T]
	toInsert Enumerable[T]
	index    int
}

func (enumerable insertAtEnumerable[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T)

	index := 0
	inserted := false
	go func() {
		defer close(ch)
		for current_a := range enumerable.values.GetEnumerator() {
			if index == enumerable.index {
				inserted = true
				YieldRange(ch, enumerable.toInsert)
			}

			ch <- current_a
			index++
		}

		if !inserted {
			YieldRange(ch, enumerable.toInsert)
		}
	}()

	return ch
}
