package task

import "net/http"

// Handler defines method for a task handler.
type Handler interface {
	// HandleGet will handle getting the task.
	HandleGet(w http.ResponseWriter, r *http.Request)

	// HandlePost will handle adding a task.
	HandlePost(w http.ResponseWriter, r *http.Request)

	// HandlePut will handle updating a task.
	HandlePut(w http.ResponseWriter, r *http.Request)

	// HandleDelete will handle deleting a task.
	HandleDelete(w http.ResponseWriter, r *http.Request)
}
