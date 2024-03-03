package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueEmpty(t *testing.T) {
	q := New[rune]()

	actual := q.Values()

	assert.Equal(t, 0, len(actual))
}

func TestQueuePushBack(t *testing.T) {
	q := New[string]()
	q.PushBack("h")
	q.PushBack("e")
	q.PushBack("l")
	q.PushBack("l")
	q.PushBack("o")

	expected := []string{"h", "e", "l", "l", "o"}
	actual := q.Values()

	assert.Equal(t, expected, actual)
}

func TestQueuePushMixed(t *testing.T) {
	q := New[string]()
	q.PushBack("l")
	q.PushBack("l")
	q.PushBack("o")
	q.PushFront("e")
	q.PushFront("h")

	expected := []string{"h", "e", "l", "l", "o"}
	actual := q.Values()

	assert.Equal(t, expected, actual)
}

func TestQueuePopBack(t *testing.T) {
	q := New[string]()
	q.PushBack("h")
	q.PushBack("e")
	q.PushBack("l")
	q.PushBack("l")
	q.PushBack("o")
	q.PopBack()

	expected := []string{"h", "e", "l", "l"}
	actual := q.Values()

	assert.Equal(t, expected, actual)
}

func TestQueuePopMixed(t *testing.T) {
	q := New[int]()
	q.PushBack(0)
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)
	q.PopFront()
	q.PopFront()
	q.PopBack()

	expected := []int{2}
	actual := q.Values()

	assert.Equal(t, expected, actual)
}
