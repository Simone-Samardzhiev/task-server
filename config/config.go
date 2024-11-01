package config

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"task-server/middleware"
	"task-server/task"
	"task-server/user"
)

// LoadEnvironmentFiles will load all the environment files from a .env file.
func LoadEnvironmentFiles(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening config file: %s", err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		values := strings.Split(line, "=")

		if len(values) != 2 {
			log.Printf("Error parsing config line: %s", line)
		}

		key := strings.TrimSpace(values[0])
		value := strings.TrimSpace(values[1])

		err = os.Setenv(key, value)
		if err != nil {
			log.Printf("Error setting environment variable: %s", err)
		}
	}

	err = file.Close()
	if err != nil {
		log.Printf("Error closing config file: %s", err)
	}
}

// CreateHandlers will create the handlers for the server.
func CreateHandlers() (user.Handler, task.Handler) {
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	jwtSecret := os.Getenv("JWT_SECRET")

	connStatement := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStatement)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewPostgresRepository(db)
	authenticator := middleware.NewJWTAuthenticator([]byte(jwtSecret), []string{"Task-App"}, "localhost")
	userService := user.NewServiceImp(userRepository, authenticator)
	userHandler := user.NewHandlerImp(userService)

	taskRepository := task.NewRepository(db)
	taskService := task.NewServiceImp(&taskRepository, authenticator)
	taskHandler := task.NewHandlerImp(&taskService)

	return userHandler, &taskHandler
}
