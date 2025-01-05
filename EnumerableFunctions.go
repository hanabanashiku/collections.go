package collections

import "errors"

func Count[T any](enumerable Enumerable[T]) int {
	if collection, ok := enumerable.(Collection[T]); ok {
		return collection.Count()
	}

	enumerator := enumerable.GetEnumerator()
	i := 0

	for enumerator.MoveNext() {
		i++
	}

	return i
}

func Any[T any](enumerable Enumerable[T], predicate func(*T) bool) bool {
	enumerator := enumerable.GetEnumerator()

	for enumerator.MoveNext() {
		if predicate(enumerator.Current()) {
			return true
		}
	}

	return false
}

func IsEmpty[T any](enumerable Enumerable[T]) bool {
	if collection, ok := enumerable.(Collection[T]); ok {
		return collection.Count() == 0
	}

	enumerator := enumerable.GetEnumerator()
	return !enumerator.MoveNext()
}

func Find[T any](enumerable Enumerable[T], predicate func(*T) bool) (*T, error) {
	enumerator := enumerable.GetEnumerator()

	for enumerator.MoveNext() {
		if predicate(enumerator.Current()) {
			return enumerator.Current(), nil
		}
	}

	return nil, errors.New("element not found")
}

func ElementAt[T any](enumerable Enumerable[T], index int) (*T, error) {
	enumerator := enumerable.GetEnumerator()
	count := 0

	for enumerator.MoveNext() {
		if count == index {
			return enumerator.Current(), nil
		}

		count++
	}

	return nil, errors.New("element not found")
}

func First[T any](enumerable Enumerable[T]) (*T, error) {
	enumerator := enumerable.GetEnumerator()

	for enumerator.MoveNext() {
		return enumerator.Current(), nil
	}

	return nil, errors.New("no elements in sequence")
}

func Last[T any](enumerable Enumerable[T]) (*T, error) {
	enumerator := enumerable.GetEnumerator()
	any := false

	for enumerator.MoveNext() {
		any = true
	}

	if !any {
		return nil, errors.New("no elements in sequence")
	}

	return enumerator.Current(), nil
}

func ForEach[T any](enumerable Enumerable[T], callback func(*T, int)) {
	enumerator := enumerable.GetEnumerator()
	index := 0

	for enumerator.MoveNext() {
		callback(enumerator.Current(), index)
		index++
	}
}

func SequenceEqual[T any](enumerable Enumerable[T], compare Enumerable[T]) bool {
	a := enumerable.GetEnumerator()
	b := compare.GetEnumerator()

	for a.MoveNext() {
		if !b.MoveNext() {
			return false
		}

		if a.Current() != b.Current() {
			return false
		}
	}

	return !b.MoveNext()
}
