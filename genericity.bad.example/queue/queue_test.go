package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue[int](5)
	for i := range make([]struct{}, 5) {
		q.Push(i)
	}

	for range make([]struct{}, 5) {
		t.Log(q.Pop())
	}
}