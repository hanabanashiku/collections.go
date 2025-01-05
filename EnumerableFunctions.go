package collections

import "errors"

func Count[T any](enumerable Enumerable[T]) int {
	if collection, ok := enumerable.(Collection[T]); ok {
		return collection.Count()
	}

	i := 0
	for _ = range enumerable.GetEnumerator() {
		i++
	}

	return i
}

func Any[T any](enumerable Enumerable[T], predicate func(*T) bool) bool {
	for current := range enumerable.GetEnumerator() {
		if predicate(current) {
			return true
		}
	}

	return false
}

func IsEmpty[T any](enumerable Enumerable[T]) bool {
	if collection, ok := enumerable.(Collection[T]); ok {
		return collection.Count() == 0
	}

	_, notEmpty := <-enumerable.GetEnumerator()
	return !notEmpty
}

func Find[T any](enumerable Enumerable[T], predicate func(*T) bool) (*T, error) {
	for current := range enumerable.GetEnumerator() {
		if predicate(current) {
			return current, nil
		}
	}

	return nil, errors.New("element not found")
}

func ElementAt[T any](enumerable Enumerable[T], index int) (*T, error) {
	count := 0

	for current := range enumerable.GetEnumerator() {
		if count == index {
			return current, nil
		}

		count++
	}

	return nil, errors.New("element not found")
}

func First[T any](enumerable Enumerable[T]) (*T, error) {

	for current := range enumerable.GetEnumerator() {
		return current, nil
	}

	return nil, errors.New("no elements in sequence")
}

func FirstBy[T any](enumerable Enumerable[T], predicate func(*T) bool) (*T, error) {
	for current := range enumerable.GetEnumerator() {
		if predicate(current) {
			return current, nil
		}
	}

	return nil, errors.New("no elements in sequence")
}

func Single[T any](enumerable Enumerable[T]) (*T, error) {
	enumerator := enumerable.GetEnumerator()
	current, ok := <-enumerator

	if !ok {
		return nil, errors.New("no matching elements in sequence")
	}

	if _, hasNext := <-enumerator; hasNext {
		return nil, errors.New("more than one matching element in sequence")
	}

	return current, nil
}

func SingleBy[T any](enumerable Enumerable[T], predicate func(*T) bool) (*T, error) {
	return Single(Where(enumerable, predicate))
}

func Last[T any](enumerable Enumerable[T]) (*T, error) {
	var last *T

	for current := range enumerable.GetEnumerator() {
		last = current
	}

	if last == nil {
		return nil, errors.New("no elements in sequence")
	}

	return last, nil
}

func LastBy[T any](enumerable Enumerable[T], predicate func(*T) bool) (*T, error) {
	return Last(Where(enumerable, predicate))
}

func ForEach[T any](enumerable Enumerable[T], callback func(*T, int)) {
	index := 0

	for current := range enumerable.GetEnumerator() {
		callback(current, index)
		index++
	}
}

func SequenceEqual[T any](enumerable Enumerable[T], compare Enumerable[T]) bool {
	a := enumerable.GetEnumerator()
	b := compare.GetEnumerator()

	for current_a := range a {
		current_b, b_open := <-b
		if !b_open {
			return false
		}

		if current_a != current_b {
			return false
		}
	}

	_, b_open := <-b
	return !b_open
}
