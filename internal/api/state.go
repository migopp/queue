package api

import "github.com/migopp/queue/internal/queue"

// Package application state
//
// Really this is only a user queue for the time being
type State struct {
	Queue queue.Queue[queue.Entry]
}

// Initialize the application state
//
// It should reset when the server does, I think...
var AppState = &State{
	Queue: queue.Queue[queue.Entry]{},
}
