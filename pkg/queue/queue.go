package queue

import "fmt"

// Queue implemented as doubly linked list
type Queue[T any] struct {
	head *item[T]
	tail *item[T]
}

// Item stored in the Queue
type item[T any] struct {
	value T
	prev  *item[T]
	next  *item[T]
}

// New constructs empty Queue
func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

// PushFront inserts a value to the front of the queue
func (q *Queue[T]) PushFront(value T) {
	item := &item[T]{
		value: value,
	}
	if q.head == nil {
		q.head = item
		q.tail = item
		return
	}

	item.next = q.head
	q.head.prev = item
	q.head = item
}

// PushBack inserts a value to the back of the queue
func (q *Queue[T]) PushBack(value T) {
	item := &item[T]{
		value: value,
	}
	if q.head == nil {
		q.head = item
		q.tail = item
		return
	}

	item.prev = q.tail
	q.tail.next = item
	q.tail = item
}

// PopFront pops a value from the front of the queue
func (q *Queue[T]) PopFront() *T {
	if q.head == nil {
		return nil
	}

	head := q.head
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		return &head.value
	}

	q.head = q.head.next
	q.head.prev = nil
	return &head.value
}

// PopBack pops a value from the back of the queue
func (q *Queue[T]) PopBack() *T {
	if q.head == nil {
		return nil
	}

	tail := q.tail
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		return &tail.value
	}

	q.tail = q.tail.prev
	q.tail.next = nil
	return &tail.value
}

// Values returns the contents of the list
func (q *Queue[T]) Values() []T {
	values := make([]T, 0)
	item := q.head
	for item != nil {
		values = append(values, item.value)
		item = item.next
	}

	return values
}

// Print prints the content of the list
func (q *Queue[T]) Print() {
	item := q.head
	for item != nil {
		fmt.Println(item.value)
		item = item.next
	}
}
