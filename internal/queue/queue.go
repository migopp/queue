package queue

import (
	"errors"

	"github.com/migopp/queue/internal/debug"
)

// A simple representation of a queue in memory
//
// > "Why? Just use a library."
// No. It took two minutes.
//
// TODO: Make thread-safe w/ sync.
type Queue[T any] struct {
	ary []T
}

// Offer to the end of the queue
func (q *Queue[T]) Offer(e T) {
	q.ary = append(q.ary, e)
}

// Poll from the front of the queue
//
// Returns the poll'd value, or the zero value and an error
// if the queue is empty
func (q *Queue[T]) Poll() (T, error) {
	var def T // initializes as the zero value
	if len(q.ary) == 0 {
		return def, errors.New("queue is empty")
	}
	def = q.ary[0]
	q.ary = q.ary[1:]
	return def, nil
}

// Transform this queue into a slice
func (q *Queue[T]) Slice() []T {
	return q.ary
}

// Debug print the contents of the queue
func (q *Queue[T]) Debug() {
	debug.Printf("| Queue Dump:\n")
	for i, item := range q.ary {
		debug.Printf("|\t%d: %v\n", i, item)
	}
}
