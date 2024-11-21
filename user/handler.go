package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"task-server/middleware"
)

// HandlerImp implements Handler.
type HandlerImp struct {
	Service Service
}

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

// HandleLogin will respond to log in requests and return a refresh token.
func (h *HandlerImp) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.handleInvalidMethod(w)
		return
	}

	var user WithoutIdUser
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Printf("Error in user-HandleImp-HandleLogin: %v", err)
		h.handleInvalidJson(w)
		return
	}

	token, err := h.Service.Login(&user)

	if errors.Is(err, ErrWrongCredentials) {
		http.Error(w, "Wrong credentials", http.StatusUnauthorized)
	} else if err != nil {
		log.Printf("Error in user-HandleImp-HandleLogin: %v", err)
		h.handleServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	_, err = fmt.Fprint(w, *token)
	if err != nil {
		log.Printf("Error in user-HandleImp-HandleLogin: %v", err)
		h.handleServerError(w)
		return
	}
}

// HandleRegister will respond to register requests and add a new user.
func (h *HandlerImp) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.handleInvalidMethod(w)
		return
	}

	var user WithoutIdUser
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Printf("Error in user-HandleImp-HandleRegister: %v", err)
		h.handleInvalidJson(w)
		return
	}

	err = h.Service.Register(&user)
	if errors.Is(err, ErrEmailInUse) {
		http.Error(w, "Email already in use", http.StatusConflict)
		return
	} else if err != nil {
		log.Printf("Error in user-HandleImp-HandleRegister: %v", err)
		h.handleServerError(w)
	}

	w.WriteHeader(http.StatusOK)
}

// HandleRefresh will respond to refresh token requests and respond with a new token.
func (h *HandlerImp) HandleRefresh(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.handleInvalidMethod(w)
		return
	}

	token, err := middleware.GetTokenFromHeader(r)
	if err != nil {
		h.handleInvalidToken(w)
		return
	}

	group, err := h.Service.RefreshTokens(&token)
	if errors.Is(err, ErrInvalidToken) {
		h.handleInvalidToken(w)
		return
	} else if err != nil {
		log.Printf("Error in user-HandleImp-HandleRefresh: %v", err)
		h.handleServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(group)
	if err != nil {
		log.Printf("Error in user-HandleImp-HandleRefresh: %v", err)
		h.handleServerError(w)
		return
	}
}

func NewHandlerImp(service Service) *HandlerImp {
	return &HandlerImp{
		Service: service,
	}
}
