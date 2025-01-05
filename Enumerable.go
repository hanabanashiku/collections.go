package collections

type Enumerable[T any] interface {
	GetEnumerator() Enumerator[T]
}

type Collection[T any] interface {
	Count() int
	Get(index int) *T
	Set(index int, item *T)
	IndexOf(item *T) int
	Contains(item *T) bool
	ToArray() []*T
	AsReadOnly() ReadOnlyCollection[T]
	GetEnumerator() Enumerator[T]
}

type ReadOnlyCollection[T any] interface {
	Count() int
	Get(index int) *T
	IndexOf(item *T) int
	GetEnumerator() Enumerator[T]
}
