package collections

func Append[T any](enumerable Enumerable[T], values Enumerable[T]) Enumerable[T] {
	return appendEnumerable[T]{
		values:   enumerable,
		toInsert: values,
		index:    -1,
	}
}

func Prepend[T any](enumerable Enumerable[T], values Enumerable[T]) Enumerable[T] {
	return appendEnumerable[T]{
		values:   enumerable,
		toInsert: values,
		index:    0,
	}
}

func Insert[T any](enumerable Enumerable[T], values Enumerable[T], index int) Enumerable[T] {
	return appendEnumerable[T]{
		values:   enumerable,
		toInsert: values,
		index:    index,
	}
}

type appendEnumerable[T any] struct {
	values   Enumerable[T]
	toInsert Enumerable[T]
	index    int
}

func (enumerable appendEnumerable[T]) GetEnumerator() Enumerator[T] {
	return appendEnumerator[T]{
		values:        enumerable.values.GetEnumerator(),
		toInsert:      enumerable.toInsert.GetEnumerator(),
		index:         enumerable.index,
		isValues:      true,
		valuesIndex:   -1,
		toInsertIndex: -1,
	}
}

type appendEnumerator[T any] struct {
	values        Enumerator[T]
	toInsert      Enumerator[T]
	isValues      bool
	index         int
	valuesIndex   int
	toInsertIndex int
}

func (enumerator appendEnumerator[T]) MoveNext() bool {
	if enumerator.isValues {
		enumerator.valuesIndex++

		if enumerator.valuesIndex == enumerator.index {
			enumerator.isValues = false
		} else if enumerator.values.MoveNext() {
			return true
		} else if enumerator.toInsertIndex > -1 {
			return false
		} else {
			enumerator.isValues = false
		}
	}

	if !enumerator.isValues {
		enumerator.toInsertIndex++

		if enumerator.toInsert.MoveNext() {
			return true
		} else if enumerator.values.MoveNext() {
			enumerator.isValues = true
			return true
		} else {
			return false
		}
	}

	return false
}

func (enumerator appendEnumerator[T]) Current() *T {
	if enumerator.isValues {
		return enumerator.values.Current()
	}

	return enumerator.toInsert.Current()
}

func (enumerator appendEnumerator[T]) Reset() {
	enumerator.values.Reset()
	enumerator.toInsert.Reset()
	enumerator.isValues = true
	enumerator.valuesIndex = -1
	enumerator.toInsertIndex = -1
}
