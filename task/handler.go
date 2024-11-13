package task

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"task-server/middleware"

	"github.com/google/uuid"
)

// HandlerImp is an implementation of Handler.
type HandlerImp struct {
	Service Service
}

// handleInvalidMethod will respond to any invalid http methods.
func (h *HandlerImp) handleInvalidMethod(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// handleInvalidJson will respond to any invalid json formats send to the server.
func (h *HandlerImp) handleInvalidJson(w http.ResponseWriter) {
	http.Error(w, "Invalid json format", http.StatusBadRequest)
}

// handleServerError will respond each time there is a server error.
func (h *HandlerImp) handleServerError(w http.ResponseWriter) {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

// handleInvalidToken will respond each time there is an invalid token.
func (h *HandlerImp) handleInvalidToken(w http.ResponseWriter) {
	http.Error(w, "Invalid token", http.StatusUnauthorized)
}

// handleInvalidPriority will respond each time there is an invalid priority.
func (h *HandlerImp) handleInvalidPriority(w http.ResponseWriter) {
	http.Error(w, "Invalid priority", http.StatusBadRequest)
}

// HandleGet will handle all get request and send all task.
func (h *HandlerImp) HandleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.handleInvalidMethod(w)
		return
	}

	token, err := middleware.GetTokenFromHeader(r)
	if err != nil {
		h.handleInvalidToken(w)
		return
	}

	tasks, err := h.Service.GetTasks(&token)
	if errors.Is(err, ErrInvalidToken) {
		h.handleInvalidToken(w)
		return
	} else if err != nil {
		log.Printf("Error in task-HandlerImp-HandleGet: %v", err)
		h.handleServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)

	if err != nil {
		log.Printf("Error in task-HandlerImp-HandleGet: %v", err)
		h.handleServerError(w)
	}
}

// HandlePost will handle post requests for adding a task.
func (h *HandlerImp) HandlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.handleInvalidMethod(w)
		return
	}

	token, err := middleware.GetTokenFromHeader(r)
	if err != nil {
		h.handleInvalidToken(w)
		return
	}

	var receivedTask NewTask
	err = json.NewDecoder(r.Body).Decode(&receivedTask)
	if err != nil {
		h.handleInvalidJson(w)
		return
	}

	newTask, err := h.Service.AddTask(&token, &receivedTask)
	if errors.Is(err, ErrInvalidToken) {
		h.handleInvalidPriority(w)
		return
	} else if err != nil {
		log.Printf("Error in task-HandlerImp-HandlePost: %v", err)
		h.handleServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newTask)
	if err != nil {
		log.Printf("Error in task-HandlerImp-HandlePost: %v", err)
		h.handleServerError(w)
	}
}

// HandlePut will handle put requests for updating a task.
func (h *HandlerImp) HandlePut(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.handleInvalidMethod(w)
		return
	}

	token, err := middleware.GetTokenFromHeader(r)
	if err != nil {
		h.handleInvalidToken(w)
		return
	}

	var receivedTask Task
	err = json.NewDecoder(r.Body).Decode(&receivedTask)
	if err != nil {
		h.handleInvalidJson(w)
		return
	}

	updatedTask, err := h.Service.UpdateTask(&token, &receivedTask)
	if errors.Is(err, ErrInvalidToken) {
		h.handleInvalidToken(w)
		return
	} else if err != nil {
		log.Printf("Error in task-HandlerImp-HandlePut: %v", err)
		h.handleServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(updatedTask)
	if err != nil {
		log.Printf("Error in task-HandlerImp-HandlePut: %v", err)
		h.handleServerError(w)
	}
}

// HandleDelete will handle delete request for deleting a task.
func (h *HandlerImp) HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.handleInvalidMethod(w)
		return
	}

	token, err := middleware.GetTokenFromHeader(r)
	if err != nil {
		h.handleInvalidToken(w)
		return
	}

	idString := r.URL.Query().Get("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteTask(&token, &id)
	if errors.Is(err, ErrInvalidToken) {
		h.handleInvalidToken(w)
		return
	} else if err != nil {
		log.Printf("Error in task-HandlerImp-HandleDelete: %v", err)
		h.handleServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func NewHandlerImp(service Service) HandlerImp {
	return HandlerImp{
		Service: service,
	}
}
