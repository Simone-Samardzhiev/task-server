package main

import (
	"log"
	"net/http"
	"task-server/config"
)

func main() {
	config.LoadEnvironmentFiles("../../config/.env")
	userHandler, taskHandler := config.CreateHandlers()

	mux := http.NewServeMux()
	mux.Handle("/users/login", http.HandlerFunc(userHandler.HandleLogin))
	mux.Handle("/users/register", http.HandlerFunc(userHandler.HandleRegister))
	mux.Handle("/users/refresh", http.HandlerFunc(userHandler.HandleRefresh))
	mux.Handle("/tasks/get", http.HandlerFunc(taskHandler.HandleGet))
	mux.Handle("/tasks/add", http.HandlerFunc(taskHandler.HandlePost))
	mux.Handle("/tasks/update", http.HandlerFunc(taskHandler.HandlePut))
	mux.Handle("/tasks/delete", http.HandlerFunc(taskHandler.HandleDelete))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
