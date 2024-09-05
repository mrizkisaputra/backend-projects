package services

import "task-tracker/entities"

type TaskServiceInterface interface {
	AppendTask(desc string) (string, error)

	UpdateTask(id int, desc string) string

	DeleteTask(id int) string

	MarkTask(status string, id int) string

	AllTasks() map[string][]entities.Task

	GetTaskByStatus(status string) map[string][]entities.Task
}
