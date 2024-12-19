package util

import "testing"

func TestPriorityQueue_Initializer(t *testing.T) {
	q := NewArrayPriorityQueue[int]()
	if q == nil {
		t.Errorf("NewArrayPriorityQueue() = nil, want not nil")
	}
}

func TestPriorityQueue_IsEmptyAtStart(t *testing.T) {
	q := NewArrayPriorityQueue[int]()
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = false, want true")
	}
}

func TestPriorityQueue_SizeAtStart(t *testing.T) {
	q := NewArrayPriorityQueue[int]()
	if q.Size() != 0 {
		t.Errorf("Size() = %d, want 0", q.Size())
	}
}

func TestPriorityQueue_SimpleInsert(t *testing.T) {
	q := NewArrayPriorityQueue[int]()
	q.Insert(1)
	if q.Size() != 1 {
		t.Errorf("Size() = %d, want 1", q.Size())
	}
	if q.IsEmpty() {
		t.Errorf("IsEmpty() = true, want false")
	}
}

func TestPriorityQueue_SimpleRemove(t *testing.T) {
	q := NewArrayPriorityQueue[int]()
	q.Insert(1)
	if q.Remove() != 1 {
		t.Errorf("Remove() = %d, want 1", q.Remove())
	}
	if q.Size() != 0 {
		t.Errorf("Size() = %d, want 0", q.Size())
	}
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = false, want true")
	}
}

func TestPriorityQueue_StrongerInsert(t *testing.T) {
	q := NewArrayPriorityQueue[int]()
	q.Insert(3)
	q.Insert(1)
	q.Insert(2)
	if q.Size() != 3 {
		t.Errorf("Size() = %d, want 3", q.Size())
	}
	if q.IsEmpty() {
		t.Errorf("IsEmpty() = true, want false")
	}
}

func TestPriorityQueue_StrongerRemove(t *testing.T) {
	q := NewArrayPriorityQueue[int]()
	q.Insert(3)
	q.Insert(1)
	q.Insert(2)
	if q.Remove() != 1 {
		t.Errorf("Remove() = %d, want 1", q.Remove())
	}
	if q.Size() != 2 {
		t.Errorf("Size() = %d, want 2", q.Size())
	}
	if q.Remove() != 2 {
		t.Errorf("Remove() = %d, want 2", q.Remove())
	}
	if q.Size() != 1 {
		t.Errorf("Size() = %d, want 1", q.Size())
	}
	if q.Remove() != 3 {
		t.Errorf("Remove() = %d, want 3", q.Remove())
	}
	if q.Size() != 0 {
		t.Errorf("Size() = %d, want 0", q.Size())
	}
	if !q.IsEmpty() {
		t.Errorf("IsEmpty() = false, want true")
	}
}
