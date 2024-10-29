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
	mux.Handle("POST users/login", http.HandlerFunc(userHandler.HandleLogin))
	mux.Handle("POST users/register", http.HandlerFunc(userHandler.HandleRegister))
	mux.Handle("GET users/refresh", http.HandlerFunc(userHandler.HandleRefresh))

	err := http.ListenAndServeTLS(":8080", "../../server-cert.pem", "../../server-key.pem", mux)
	if err != nil {
		log.Fatal(err)
	}
}
