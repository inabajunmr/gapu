package ufap

import (
	"fmt"
	"testing"
)

func TestQueueSingleValue(t *testing.T) {
	q := Queue{}
	if !q.IsEmpty() {
		t.Fatalf("init queue is not empty")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Fatalf("queue is empty")
	}
	q.Enqueue(2)
	q.Enqueue(3)
	assertDequeue(&q, 1, t)
	if q.IsEmpty() {
		t.Fatalf("queue is empty")
	}
	assertDequeue(&q, 2, t)
	assertDequeue(&q, 3, t)
	fmt.Println(len(q.val))
	if !q.IsEmpty() {
		t.Fatalf("queue is not empty")
	}

}

func assertDequeue(queue *Queue, expected int, t *testing.T) {
	val, ok := queue.Dequeue().(int)
	if !ok {
		t.Fatalf("dequeue error")
	}
	if val != expected {
		t.Fatalf("dequeue error")
	}
}
