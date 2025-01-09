package collections

type iComparable interface {
	comparable
}

func Aggregate[TSource any, TAccumulate any](
	source Enumerable[TSource],
	accumulator func(TAccumulate, *TSource) TAccumulate,
	initialValue TAccumulate) TAccumulate {
	if source == nil {
		panic("source must not be nil")
	}
	if accumulator == nil {
		panic("accumulator must not be nil")
	}

	result := initialValue

	for currentValue := range source.GetEnumerator() {
		result = accumulator(result, currentValue)
	}

	return result
}

func AggregateBy[TSource any, TKey iComparable, TAccumulate any](
	source Enumerable[TSource],
	keySelector func(*TSource) TKey,
	accumulator func(TAccumulate, *TSource) TAccumulate,
	initialSelector func(TKey) TAccumulate) Enumerable[KeyValuePair[TKey, TAccumulate]] {
	result := make(map[TKey]TAccumulate)

	for currentValue := range source.GetEnumerator() {
		key := keySelector(currentValue)
		list, ok := result[key]

		if !ok {
			list = initialSelector(key)
		}

		result[key] = accumulator(list, currentValue)
	}

	return aggregateByEnumerable[TKey, TAccumulate]{result}
}

type aggregateByEnumerable[TKey iComparable, TAccumulate any] struct {
	results map[TKey]TAccumulate
}

func (enumerable aggregateByEnumerable[TKey, TAccumulate]) GetEnumerator() Enumerator[KeyValuePair[TKey, TAccumulate]] {
	ch := make(chan *KeyValuePair[TKey, TAccumulate])

	go func() {
		defer close(ch)
		for key, value := range enumerable.results {
			item := KeyValuePair[TKey, TAccumulate]{key, value}
			ch <- &item
		}

	}()

	return ch
}
