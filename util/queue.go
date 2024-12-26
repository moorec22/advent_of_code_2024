package util

type Queue[T any] interface {
	// Insert inserts an item into the queue
	Insert(T)
	// Remove removes the top item from the queue.
	Remove() T
	// Size returns the number of items in the queue
	Size() int
	// IsEmpty returns true if the queue is empty
	IsEmpty() bool
}

type ArrayQueue[T any] struct {
	arr  []T
	size int
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{make([]T, 0), 0}
}

func (q *ArrayQueue[T]) Insert(t T) {
	if len(q.arr) == q.size {
		q.arr = append(q.arr, t)
	} else {
		q.arr[q.size] = t
		q.size++
	}
}

func (q *ArrayQueue[T]) Remove() T {
	el := q.arr[q.size-1]
	q.size--
	return el
}

func (q *ArrayQueue[T]) Size() int {
	return q.size
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.size == 0
}
