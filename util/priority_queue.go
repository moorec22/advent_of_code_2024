package util

import "cmp"

type PriorityQueue[T cmp.Ordered] interface {
	// Insert inserts an item into the priority queue
	Insert(T)
	// Remove removes the top item from the priority queue.
	Remove() T
	// Size returns the number of items in the priority queue
	Size() int
	// IsEmpty returns true if the priority queue is empty
	IsEmpty() bool
}

type ArrayPriorityQueue[T cmp.Ordered] struct {
	data []T
	size int
}

func NewArrayPriorityQueue[T cmp.Ordered]() *ArrayPriorityQueue[T] {
	return &ArrayPriorityQueue[T]{make([]T, 0), 0}
}

func (q *ArrayPriorityQueue[T]) Insert(item T) {
	q.data = append(q.data, item)
	q.size++
	q.percolateUp()
}

func (q *ArrayPriorityQueue[T]) Remove() T {
	if q.size == 0 {
		panic("empty queue")
	}
	top := q.data[0]
	q.data[0] = q.data[q.size-1]
	q.size--
	q.percolateDown()
	return top
}

func (q *ArrayPriorityQueue[T]) Size() int {
	return q.size
}

func (q *ArrayPriorityQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *ArrayPriorityQueue[T]) percolateUp() {
	nodeIndex := q.size - 1
	parentIndex := q.parentIndex(nodeIndex)
	for nodeIndex > 0 && q.data[nodeIndex] < q.data[parentIndex] {
		q.swap(nodeIndex, parentIndex)
		nodeIndex = parentIndex
		parentIndex = q.parentIndex(nodeIndex)
	}
}

func (q *ArrayPriorityQueue[T]) percolateDown() {
	nodeIndex := 0
	leftChild, rightChild := q.childrenIndices(nodeIndex)
	for {
		smallestChild := -1
		if leftChild < q.size && q.data[leftChild] < q.data[nodeIndex] {
			smallestChild = leftChild
		}
		if rightChild < q.size && q.data[rightChild] < q.data[nodeIndex] && q.data[rightChild] < q.data[leftChild] {
			smallestChild = rightChild
		}
		if smallestChild == -1 {
			break
		} else {
			q.swap(nodeIndex, smallestChild)
			nodeIndex = smallestChild
		}
		leftChild, rightChild = q.childrenIndices(nodeIndex)
	}
}

func (q *ArrayPriorityQueue[T]) swap(one, two int) {
	temp := q.data[one]
	q.data[one] = q.data[two]
	q.data[two] = temp
}

func (q *ArrayPriorityQueue[T]) parentIndex(index int) int {
	return index / 2
}

func (q *ArrayPriorityQueue[T]) childrenIndices(index int) (int, int) {
	return 2 * index, 2*index + 1
}
