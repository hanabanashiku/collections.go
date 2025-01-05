package collections

import "errors"

type Stack[T any] struct {
	items []*T
	count int
}

func NewStack[T any]() Stack[T] {
	items := make([]*T, 0, defaultCapacity)
	return Stack[T]{
		items: items,
		count: 0,
	}
}

func NewStackWithCapacity[T any](capacity int) Stack[T] {
	if capacity < 0 {
		panic("capacity out of range")
	}

	items := make([]*T, 0, capacity)
	return Stack[T]{
		items: items,
		count: 0,
	}
}

func NewStackFromEnumerable[T any](enumerable Enumerable[T]) Stack[T] {
	items := ToList(enumerable)

	return Stack[T]{
		items: items.ToArray(),
		count: items.Count(),
	}
}

func NewStackFromArray[T any](array []*T) Stack[T] {
	return Stack[T]{
		items: array,
		count: len(array),
	}
}

func (stack Stack[T]) Peek() (*T, error) {
	if stack.count <= 0 {
		return nil, errors.New("stack is empty")
	}

	return stack.items[stack.count-1], nil
}

func (stack Stack[T]) Pop() (*T, error) {
	if stack.count <= 0 {
		return nil, errors.New("stack is empty")
	}

	value := stack.items[stack.count-1]
	stack.items = stack.items[:stack.count-1]
	stack.count--
	return value, nil
}

func (stack Stack[T]) Push(element *T) {
	stack.count++
	if stack.GetCapacity() <= stack.count {
		growStack(&stack, stack.count-stack.GetCapacity())
	}

	stack.items[stack.count-1] = element
}

func (stack Stack[T]) Clear() {
	stack.count = 0
	stack.items = make([]*T, 0, len(stack.items))
}

func (stack Stack[T]) GetCapacity() int {
	return len(stack.items)
}

func (stack Stack[T]) Count() int {
	return stack.count
}

func (stack Stack[T]) ToArray() []*T {
	array := stack.items
	return array[:stack.count]
}

func (stack Stack[T]) GetEnumerator() Enumerator[T] {
	ch := make(chan *T, stack.count)

	go func() {
		for i := 0; i < stack.count; i++ {
			ch <- stack.items[i]
		}
		close(ch)
	}()

	return ch
}

func growStack[T any](stack *Stack[T], increaseBy int) {
	newItems := make([]*T, 0, stack.GetCapacity()+increaseBy)
	copy(newItems, stack.items)
	stack.items = newItems
}
