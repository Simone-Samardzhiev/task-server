package task

import (
	"database/sql"
	"log"
	"task-server/middleware"

	"github.com/google/uuid"
)

// ServiceImp is an implementation of Service.
type ServiceImp struct {
	Repository    Repository
	Authenticator middleware.Authenticator
}

// GetTasks will return a slice of all tasks that belongs to a user.
func (s *ServiceImp) GetTasks(tokenString *string) ([]Task, error) {
	id, err := s.Authenticator.CheckAccessToken(tokenString)
	if err != nil {
		log.Printf("Error in task-ServiceImp-GetTasks: %v", err)
		return nil, ErrInvalidToken
	}

	tasks, err := s.Repository.GetTasks(id)
	if err != nil {
		log.Printf("Error in task-ServiceImp-GetTasks: %v", err)
		return nil, err
	}
	return tasks, nil
}

// AddTask will add a new task to a user.
func (s *ServiceImp) AddTask(tokenString *string, newTask *NewTask) (*Task, error) {
	id, err := s.Authenticator.CheckAccessToken(tokenString)
	if err != nil {
		log.Printf("Error in task-ServiceImp-AddTask: %v", err)
		return nil, ErrInvalidToken
	}

	ok, err := s.Repository.CheckPriority(&newTask.Priority)
	if err != nil {
		log.Printf("Error in task-ServiceImp-AddTask: %v", err)
		return nil, err
	}
	if !ok {
		return nil, ErrInvalidPriority
	}

	task := &Task{
		Id:            uuid.New(),
		Name:          newTask.Name,
		Description:   newTask.Description,
		Priority:      newTask.Priority,
		DueDate:       newTask.DueDate,
		DateDeleted:   NullTime{sql.NullTime{Valid: false}},
		DateCompleted: NullTime{sql.NullTime{Valid: false}},
	}
	err = s.Repository.AddTask(task, id)
	if err != nil {
		log.Printf("Error in task-ServiceImp-AddTask: %v", err)
		return nil, err
	}

	return task, nil
}

// UpdateTask will update an existing task information.
func (s *ServiceImp) UpdateTask(stringToken *string, task *Task) (*Task, error) {
	_, err := s.Authenticator.CheckAccessToken(stringToken)
	if err != nil {
		log.Printf("Error in task-ServiceImp-UpdateTask: %v", err)
		return nil, ErrInvalidToken
	}

	ok, err := s.Repository.CheckPriority(&task.Priority)
	if err != nil {
		log.Printf("Error in task-ServiceImp-UpdateTask: %v", err)
	}

	if !ok {
		return nil, ErrInvalidPriority
	}

	err = s.Repository.UpdateTask(task)
	if err != nil {
		log.Printf("Error in task-ServiceImp-UpdateTask: %v", err)
		return nil, err
	}
	return task, nil
}

// DeleteTask will delete a existing task.
func (s *ServiceImp) DeleteTask(tokenString *string, uuid *uuid.UUID) error {
	_, err := s.Authenticator.CheckAccessToken(tokenString)
	if err != nil {
		log.Printf("Error in task-ServiceImp-DeleteTask: %v", err)
		return ErrInvalidToken
	}

	err = s.Repository.DeleteTask(uuid)
	if err != nil {
		log.Printf("Error in task-ServiceImp-DeleteTask: %v", err)
		return err
	}

	return nil
}

// NewServiceImp will create a new service with a authenticator and repository.
func NewServiceImp(repository Repository, authenticator middleware.Authenticator) ServiceImp {
	return ServiceImp{
		Repository:    repository,
		Authenticator: authenticator,
	}
}
