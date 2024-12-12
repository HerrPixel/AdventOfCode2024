package Tools

import "fmt"

type Queue[T any] struct {
	data []T
}

func Size[T any](q Queue[T]) int {
	return len(q.data)
}

func IsEmpty[T any](q Queue[T]) bool {
	return Size(q) == 0
}

func Enqueue[T any](q Queue[T], o T) Queue[T] {
	q.data = append(q.data, o)

	return q
}

func Dequeue[T any](q Queue[T]) (T, Queue[T], error) {
	if Size(q) == 0 {
		var zero T
		return zero, q, fmt.Errorf("queue is empty")
	}

	result := q.data[0]
	q.data = q.data[1:]

	return result, q, nil
}
