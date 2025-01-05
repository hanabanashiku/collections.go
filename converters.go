package collections

func ToArray[T any](enumerable Enumerable[T]) []*T {
	if collection, ok := enumerable.(Collection[T]); ok {
		return collection.ToArray()
	}

	enumerator := enumerable.GetEnumerator()
	count := 0
	buffer := make([]*T, 0, 4)

	for enumerator.MoveNext() {
		if count == len(buffer) {
			newBuffer := make([]*T, 0, len(buffer)*2)
			copy(newBuffer, buffer)
			buffer = newBuffer
		}

		count++
		buffer[count] = enumerator.Current()
	}

	array := make([]*T, 0, count)
	copy(array, buffer[:count])
	return array
}

func ToList[T any](enumerable Enumerable[T]) List[T] {
	if list, ok := enumerable.(List[T]); ok {
		return list
	}

	items := ToArray(enumerable)
	return List[T]{
		items: items,
		count: len(items),
	}
}
