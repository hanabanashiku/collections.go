package collections

func Range(start, count int) Enumerable[int] {
	if count < 0 {
		panic("count out of range")
	}

	return rangeEnumerable{
		start: start,
		count: count,
	}
}

func Repeat[T any](element *T, count int) Enumerable[T] {
	if count < 0 {
		panic("count out of range")
	}

	return repeatEnumerable[T]{
		element: element,
		count:   count,
	}

}

type rangeEnumerable struct {
	start, count int
}

func (enumerable rangeEnumerable) GetEnumerator() Enumerator[int] {
	ch := make(chan *int, enumerable.count)

	go func() {
		for current := enumerable.start; current < enumerable.start+enumerable.count; current++ {
			ch <- &current
		}
		close(ch)
	}()

	return ch
}

type repeatEnumerable[T any] struct {
	element *T
	count   int
}

func (enumerable repeatEnumerable[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T, enumerable.count)

	go func() {
		for i := 0; i < enumerable.count; i++ {
			ch <- enumerable.element
		}
		close(ch)
	}()

	return ch
}
