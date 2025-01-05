package collections

import (
	"errors"
)

const defaultCapacity = 4

type List[T any] struct {
	items []*T
	count int
}

func NewList[T any]() List[T] {
	return List[T]{
		items: make([]*T, 0, defaultCapacity),
		count: 0,
	}
}

func NewListWithCapacity[T any](capacity int) List[T] {
	return List[T]{
		items: make([]*T, 0, capacity),
		count: 0,
	}
}

func NewListFromArray[T any](array []T) List[T] {
	items := make([]*T, len(array))

	for i, item := range array {
		items[i] = &item
	}

	return List[T]{
		items: items,
		count: len(array),
	}
}

func (list List[T]) Count() int {
	return list.count
}

func (list List[T]) GetCapacity() int {
	return len(list.items)
}

func (list List[T]) ToArray() []*T {
	array := list.items
	return array[:list.count]
}

func (list List[T]) Get(i int) *T {
	assertSize(&list, i)
	return list.items[i]
}

func (list List[T]) Set(i int, item *T) {
	assertSize(&list, i)
	list.items[i] = item
}

func (list List[T]) IndexOf(item *T) int {
	index := -1

	for i, x := range list.items {
		if x == item {
			index = i
			break
		}
	}

	return index
}

func (list List[T]) Contains(item *T) bool {
	return list.IndexOf(item) >= 0
}

func (list List[T]) Add(item *T) {
	newCount := list.count + 1
	if newCount >= len(list.items) {
		growList(&list, 1)
	}

	list.items[newCount-1] = item
	list.count = newCount
}

func (list List[T]) AddRange(enumerable Enumerable[T]) {
	if collection, ok := enumerable.(Collection[T]); ok {
		growList(&list, collection.Count())
		copy(list.items[list.count:], collection.ToArray())
		list.count = list.count + collection.Count()
		return
	}

	enumerator := enumerable.GetEnumerator()
	for current := range enumerator {
		list.Add(current)
	}
}

func (list List[T]) RemoveLast() {
	_ = list.RemoveAt(list.Count() - 1)
}

func (list List[T]) RemoveAt(index int) error {
	error := checkSize(&list, index)
	if error != nil {
		return error
	}

	list.count--
	copy(list.items[index+1:], list.items[index+2:])
	list.items[len(list.items)-1] = nil
	return nil
}

func (list List[T]) RemoveRange(index, count int) error {
	if error := checkSize(&list, index); error != nil {
		return error
	}
	if count < 0 || list.count-index < count {
		return errors.New("count out of range")
	}

	list.count = list.count - count
	copy(list.items[index:], list.items[index+count+1:])

	return nil
}

func (list List[T]) Remove(item *T) bool {
	index := list.IndexOf(item)

	if index < 0 {
		return false
	}

	list.RemoveAt(index)
	return true
}

func (list List[T]) Clear() {
	if list.count == 0 {
		return
	}

	list.count = 0
	list.items = make([]*T, 0, len(list.items))
}

func (list List[T]) Reverse() {
	if list.count == 0 {
		return
	}

	list.ReverseRange(0, list.count)
}

func (list List[T]) ReverseRange(index, count int) {
	assertSize(&list, index)

	if count < 0 || list.count-index < count {
		panic("count out of range")
	}

	left := index
	right := index + count

	for left < right {
		list.items[left], list.items[right] = list.items[right], list.items[left]
		left++
		right--
	}
}

func (list List[T]) EnsureCapacity(capacity int) {
	if len(list.items) >= capacity {
		return
	}

	growList(&list, capacity-len(list.items))
}

func (list List[T]) TrimExcess() {
	if len(list.items) <= list.count {
		return
	}

	items := make([]*T, 0, list.count)
	copy(items, list.items[:list.count])
	list.items = items
}

func (list List[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T)

	go func() {
		for _, v := range list.items {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func growList[T any](list *List[T], increaseBy int) {
	newItems := make([]*T, 0, list.GetCapacity()+increaseBy)
	copy(newItems, list.items)
	list.items = newItems
}

func assertSize[T any](list *List[T], index int) {
	error := checkSize(list, index)

	if error != nil {
		panic(error.Error())
	}
}

func checkSize[T any](list *List[T], index int) error {
	if index >= 0 && index < list.Count() {
		return nil
	}

	return errors.New("index out of range")
}
