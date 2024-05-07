package iterator

type Iterator[I ~[]T, T any] struct {
	data  I
	index int
}

func NewIterator[I ~[]T, T any](data I) *Iterator[I, T] {
	return &Iterator[I, T]{
		data:  data,
		index: -1,
	}
}

func (i *Iterator[I, T]) Next() bool {
	i.index++
	return i.index < len(i.data)
}

func (i *Iterator[I, T]) Value() T {
	return i.data[i.index]
}

func (i *Iterator[I, T]) Filter(fn func(T) bool) *Iterator[I, T] {
	var newData I
	for i.Next() {
		val := i.Value()
		if fn(val) {
			newData = append(newData, val)
		}
	}

	return NewIterator(newData)
}