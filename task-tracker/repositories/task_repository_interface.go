package repositories

import . "task-tracker/entities"

type TaskRepositoryInterface interface {
	// CreateFile to creating file JSON
	CreateFile(filename string) (string, error)

	// WriteFile to write task in a file JSON
	WriteFile(filename string, tasks map[string][]Task) error

	// ReadFile to read all of task on file JSON. this method for append,delete,remove task
	ReadFile(filename string) (map[string][]Task, error)

	AppendTask(task Task, filename string) (bool, error)

	UpdateTask(id int, desc string, filename string) (bool, error)

	DeleteTask(id int, filename string) (bool, error)

	MarkTask(status string, id int, filename string) (bool, error)

	GetAllTask(filename string) (map[string][]Task, error)

	GetTaskByStatus(status string, filename string) (map[string][]Task, error)
}
