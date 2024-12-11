package api

import (
	"fmt"
	"net/http"

	"github.com/migopp/queue/internal/debug"
)

// Representation of the running server in memory
//
// I guess it's better not to hard-code things,
// I may want the ability to configure these things
// in the future
type Server struct {
	IP   string
	Port uint16
}

// Starts running the designated server
//
// If it fails, then we exit with an error,
// either on start-up or in the middle of handling --
// We can't know...
func (s *Server) Run() error {
	// For now, we just use the default `mux` in stdlib
	//
	// We could use something more advanced, but this
	// will do for now, since our API is not anything
	// revolutionary
	//
	// The enhancements from 1.22 are more than enough:
	// https://go.dev/blog/routing-enhancements
	http.HandleFunc("GET /", home)
	http.HandleFunc("POST /to-add", add)
	http.HandleFunc("GET /to-rem", rem)

	// Log where we are serving
	serveAddr := fmt.Sprintf("%s:%d", s.IP, s.Port)
	debug.Printf("| Starting server at %s\n", serveAddr)

	// Serve
	//
	// `ListenAndServe` returns an error, so we will
	// just bubble it up when it happens
	return http.ListenAndServe(serveAddr, nil)
}
