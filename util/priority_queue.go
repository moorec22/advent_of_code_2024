package util

// An implementation to be used for structs to be compared.
type StandardComparable[T any] interface {
	Compare(T) int
}

type PriorityQueue[T StandardComparable[T]] interface {
	// Insert inserts an item into the priority queue
	Insert(T)
	// Remove removes the top item from the priority queue.
	Remove() T
	// Size returns the number of items in the priority queue
	Size() int
	// IsEmpty returns true if the priority queue is empty
	IsEmpty() bool
}

type ArrayPriorityQueue[T StandardComparable[T]] struct {
	data []T
	size int
}

func NewArrayPriorityQueue[T StandardComparable[T]]() *ArrayPriorityQueue[T] {
	return &ArrayPriorityQueue[T]{make([]T, 0), 0}
}

func (q *ArrayPriorityQueue[T]) Insert(item T) {
	if q.size == len(q.data) {
		q.data = append(q.data, item)
	} else {
		q.data[q.size] = item
	}
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
	for nodeIndex > 0 && q.data[nodeIndex].Compare(q.data[parentIndex]) < 0 {
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
		if leftChild < q.size && q.data[leftChild].Compare(q.data[nodeIndex]) < 0 {
			smallestChild = leftChild
		}
		if rightChild < q.size && q.data[rightChild].Compare(q.data[nodeIndex]) < 0 && q.data[rightChild].Compare(q.data[leftChild]) < 0 {
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
