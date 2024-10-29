package user

import "net/http"

// Handler defines methods for a user handler.
type Handler interface {
	// HandleLogin will handle login requests.
	HandleLogin(w http.ResponseWriter, r *http.Request)

	// HandleRegister will handle register requests.
	HandleRegister(w http.ResponseWriter, r *http.Request)

	// HandleRefresh will handle refresh requests/
	HandleRefresh(w http.ResponseWriter, r *http.Request)
}
