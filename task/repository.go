package task

import (
	"database/sql"
	"github.com/google/uuid"
	"log"
)

// PostgresRepository is an implementation of Repository.
type PostgresRepository struct {
	database *sql.DB
}

// countTasks will count all the tasks of a user.
func (r *PostgresRepository) countTasks(id *uuid.UUID) (int64, error) {
	query := "SELECT COUNT(id) FROM tasks WHERE user_id  = $1"
	log.Printf("Executing query in task-PostgresRepository-countTasks: %s | Parameters %s", query, id.String())
	row := r.database.QueryRow(query, *id)

	var count int64
	err := row.Scan(&count)
	if err != nil {
		log.Printf("Error in task-PostgresRepsitory-countTasks: %v", err)
		return 0, err
	}
	return count, nil
}

// GetTasks will get all tasks of a user.
func (r *PostgresRepository) GetTasks(id *uuid.UUID) ([]Task, error) {
	query := "SELECT * FROM tasks WHERE user_id = $1"
	log.Printf("Executing query in task-PostgresRepository-GetTasks: %s | Parameters %s", query, id.String())

	rows, err := r.database.Query(query, *id)
	if err != nil {
		log.Printf("Error in task-PostgresRepsitory-GetTasks: %v", err)
		return nil, err
	}

	count, err := r.countTasks(id)
	if err != nil {
		log.Printf("Error in task-PostgresRepsitory-GetTasks: %v", err)
		return nil, err
	}

	tasks := make([]Task, 0, count)

	for rows.Next() {
		var task Task
		err = rows.Scan(&task.Id, &task.Name, &task.Description, &task.DueDate, &task.DateCompleted, &task.DateDeleted)
		if err != nil {
			log.Printf("Error in task-PostgresRepsitory-GetTasks: %v", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

// CheckPriority will check if a priority is found in the database.
func (r *PostgresRepository) CheckPriority(priority *int64) (bool, error) {
	query := "SELECT COUNT(id) FROM tasks WHERE priority = $1"
	log.Printf("Executing query in task-PostgresRepository-CheckPriority: %s | Parameters %d", query, priority)

	row := r.database.QueryRow(query, priority)
	var count int64

	err := row.Scan(&count)
	if err != nil {
		log.Printf("Error in task-PostgresRepsitory-CheckPriority: %v", err)
		return false, err
	}

	return count > 0, nil
}

// AddTask will add a new task to a user.
func (r *PostgresRepository) AddTask(task *Task, id *uuid.UUID) error {
	query := "INSERT INTO tasks(id, name, description, priority, due_date, date_completed, date_deleted, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	log.Printf("Executing query in task-PostgresRepository-AddTask: %s | Parameters %s, %s, %s, %d, %s, %s, %s, %s", query, task.Id, task.Name, task.Description, task.Priority, task.DueDate, task.DateCompleted, task.DateDeleted, id)

	_, err := r.database.Exec(query, task.Id, task.Name, task.Description, task.Priority, task.DueDate, task.DateCompleted, task.DateDeleted, *id)
	if err != nil {
		log.Printf("Error in task-PostgresRepsitory-AddTasks: %v", err)
	}
	return err
}

// UpdateTask will update an existing task.
func (r *PostgresRepository) UpdateTask(task *Task) error {
	query := "UPDATE tasks SET name = $1, description = $2, priority = $3, due_date = $4, date_completed = $5, date_deleted = $6 WHERE id = $7"
	log.Printf("Execting query in task-PostgresRepository-UpdateTask: %s | Parameters %s, %s, %d, %s, %s, %s, %s, ", query, task.Name, task.Description, task.Priority, task.DueDate, task.DateCompleted, task.DateDeleted, task.Id)

	_, err := r.database.Exec(query, task.Name, task.Description, task.Priority, task.DueDate, task.DateCompleted, task.DateDeleted, task.Id)
	if err != nil {
		log.Printf("Error in task-PostgresRepsitory-UpdateTasks: %v", err)
	}
	return err
}

// DeleteTask will delete an existing task.
func (r *PostgresRepository) DeleteTask(id *uuid.UUID) error {
	query := "DELETE FROM tasks WHERE id = $1"
	log.Printf("Executing query in task-PostgresRepository-DeleteTask: %s | Parameters %s", query, id)

	_, err := r.database.Exec(query, *id)
	if err != nil {
		log.Printf("Error in task-PostgresRepsitory-DeleteTasks: %v", err)
	}
	return err
}

// NewRepository will create a PostgresRepository.
func NewRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		database: db,
	}
}
