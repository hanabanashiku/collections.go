package collections

func isCollection[T any](enumerable Enumerable[T]) (bool, Collection[T]) {
	if collection, ok := enumerable.(Collection[T]); ok {
		return true, collection
	}

	return false, nil
}
