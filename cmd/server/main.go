package main

import (
	"log"
	"net/http"
	"task-server/config"
)

func main() {
	config.LoadEnvironmentFiles("../../config/.env")
	userHandler := config.CreateHandlers()

	mux := http.NewServeMux()
	mux.Handle("/users/login", http.HandlerFunc(userHandler.HandleLogin))
	mux.Handle("/users/register", http.HandlerFunc(userHandler.HandleRegister))
	mux.Handle("/users/refresh", http.HandlerFunc(userHandler.HandleRefresh))

	err := http.ListenAndServeTLS(":8080", "../../server-cert.pem", "../../server-key.pem", mux)
	if err != nil {
		log.Fatal(err)
	}
}
