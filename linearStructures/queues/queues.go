package queues

import "fmt"

type queues struct {
	index int
	end   int
	size  int
	queue []any
}

func NewQueue() *queues {
	return &queues{
		index: 0,
		end:   0,
		queue: make([]any, 3),
	}
}

func (q *queues) Enqueue(value any) (err error) {
	if q.size == len(q.queue) {
		err = fmt.Errorf("Queuer overflow")
		return
	}

	q.queue[q.end] = value
	q.end = (q.end + 1) % len(q.queue)
	q.size++

	return
}

func (q *queues) Dequeue() (value any, err error) {
	if q.size == 0 {
		err = fmt.Errorf("Queuer is empty")
		return
	}

	value = q.queue[q.index]
	q.index = (q.index + 1) % len(q.queue)
	q.size--

	return
}

func (q *queues) Peek() (value any, err error) {
	if q.size == 0 {
		err = fmt.Errorf("Queuer is empty")
		return
	}

	value = q.queue[q.index]
	return
}

func (q *queues) Size() int {
	return q.size
}
