package collections

func Skip[T any](source Enumerable[T], count int) Enumerable[T] {
	return skipEnumerable[T]{
		source: source,
		count:  count,
	}
}

func SkipWhile[T any](source Enumerable[T], predicate func(*T, int) bool) Enumerable[T] {
	return skipWhileEnumerable[T]{
		source:    source,
		predicate: predicate,
	}
}

func SkipLast[T any](source Enumerable[T], count int) Enumerable[T] {
	return skipLastEnumerable[T]{
		source: source,
		count:  count,
	}
}

type skipEnumerable[T any] struct {
	source Enumerable[T]
	count  int
}

func (enumerable skipEnumerable[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T)

	go func() {
		i := 0
		for current := range enumerable.source.GetEnumerator() {
			if i < enumerable.count {
				continue
			}

			ch <- current
			i++
		}

		close(ch)
	}()

	return ch
}

type skipWhileEnumerable[T any] struct {
	source    Enumerable[T]
	predicate func(*T, int) bool
}

func (enumerable skipWhileEnumerable[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T)

	go func() {
		i := 0
		skipping := true
		for current := range enumerable.source.GetEnumerator() {
			if skipping && !enumerable.predicate(current, i) {
				skipping = false
			}

			if !skipping {
				ch <- current
			}
		}

		close(ch)
	}()

	return ch
}

type skipLastEnumerable[T any] struct {
	source Enumerable[T]
	count  int
}

func (enumerable skipLastEnumerable[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T)

	go func() {
		stack := NewStack[T]()

		for current := range enumerable.source.GetEnumerator() {
			stack.Push(current)
		}

		for i := 0; i < enumerable.count; i++ {
			_, error := stack.Pop()

			if error != nil {
				close(ch)
				return
			}
		}

		for _, item := range stack.ToArray() {
			ch <- item
		}

		close(ch)
	}()

	return ch
}
