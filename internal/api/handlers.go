package api

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/migopp/queue/internal/debug"
	"github.com/migopp/queue/internal/queue"
)

// Serve the initial page
func home(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving queue.html\n")

	// Templates live in `/web/templates`
	tmplPath := filepath.Join("web", "templates", "queue.html")

	// Parse the template with name `tmplName`
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		debug.Printf("| Error loading template with name `queue.html`\n")
		http.Error(w, fmt.Sprintf("ERR LOADING TEMPLATE %v", err), http.StatusInternalServerError)
		return
	}

	// Actually load the template by writing to the response
	//
	// Looks like we can send dynamic data as well:
	// https://pkg.go.dev/html/template#Template.Execute
	//
	// I don't really know how this works for now so it's a later problem
	err = tmpl.Execute(w, AppState.Queue.Slice())
	if err != nil {
		debug.Printf("| Error executing template with name `queue.html`\n")
		http.Error(w, fmt.Sprintf("ERR EXECUTING TEMPLATE %v", err), http.StatusInternalServerError)
		return
	}
}

// Serve addition
func add(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving add request\n")

	// Get the form data -> convert
	//
	// We have it set to trigger a POST request to
	// `/to-add` with id `user-id` and req length `req-len`
	r.ParseForm()
	id := r.FormValue("user-id")
	rawReqLen := r.FormValue("req-len")
	reqLen, err := queue.ToLength(rawReqLen)
	if err != nil {
		debug.Printf("| Error retrieving `req-len`\n")
		http.Error(w, fmt.Sprintf("ERR RETRIEVING FORM DATA %v", err), http.StatusInternalServerError)
		return
	}

	// Insert into the queue
	AppState.Queue.Offer(queue.Entry{
		ID:     id,
		Length: reqLen,
	})

	// Debug log
	AppState.Queue.Debug()
	debug.Printf("| Form value: [user-id: %s, req-len: %s]\n", id, rawReqLen)

	// No response is needed.
	// We don't need to update the DOM.
}

// Serve removal
func rem(w http.ResponseWriter, r *http.Request) {
	debug.Printf("| Serving rem request\n")

	// Remove from the queue
	//
	// For now, we just support polling from the head
	_, err := AppState.Queue.Poll()
	if err != nil {
		debug.Printf("| Error in polling from empty queue\n")
		http.Error(w, fmt.Sprintf("ERR POLLING FROM QUEUE %v", err), http.StatusInternalServerError)
		return
	}

	// No response is needed.
	// We don't need to update the DOM.
}
