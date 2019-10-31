package ufap

// Queue is Queue
type Queue struct {
	val []interface{}
}

// Enqueue value
func (queue *Queue) Enqueue(val interface{}) {
	queue.val = append(queue.val, val)
}

// Dequeue value
func (queue *Queue) Dequeue() interface{} {
	res := queue.val[0]
	queue.val = queue.val[1:len(queue.val)]
	return res
}

// IsEmpty return true when queue is empty
func (queue *Queue) IsEmpty() bool {
	return len(queue.val) == 0
}
