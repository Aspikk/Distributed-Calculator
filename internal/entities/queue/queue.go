package queue

type node[T any] struct {
	value    T
	previous *node[T]
}

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
	}
}

func (q *Queue[T]) Enqueue(value T) {
	new := &node[T]{
		value: value,
	}
	if q.head == nil && q.tail == nil {
		q.head = new
		new.previous = nil
	} else {
		q.tail.previous = new
		new.previous = nil
	}

	q.tail = new
}

func (q *Queue[T]) Dequeue() (bool, T) {
	if q.head == nil {
		var null T
		return false, null
	}

	result := q.head.value
	q.head = q.head.previous

	return true, result
}
