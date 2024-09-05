package services

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"task-tracker/entities"
	"task-tracker/helper"
	. "task-tracker/repositories"
	"time"
)

var filename = "tasks.json"

type taskService struct {
	Log            *logrus.Logger
	TaskRepository TaskRepositoryInterface
}

func NewTaskService(log *logrus.Logger, taskRepository TaskRepositoryInterface) TaskServiceInterface {
	_, err := taskRepository.CreateFile(filename)
	helper.LogIfError(log, err)
	return &taskService{
		Log:            log,
		TaskRepository: taskRepository,
	}
}

func generateId() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int()
}

func (ts *taskService) AppendTask(desc string) (string, error) {
	timeStr := time.Now().Local().Format(time.RFC3339)
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	helper.LogIfError(ts.Log, err)

	task := entities.Task{
		Id:          generateId(),
		Description: desc,
		Status:      helper.TODO,
		CreatedAt:   parsedTime,
		UpdatedAt:   parsedTime,
	}
	ok, errAppendTask := ts.TaskRepository.AppendTask(task, filename)
	helper.LogIfError(ts.Log, errAppendTask)
	if ok {
		return fmt.Sprintf("Task added successfully (ID: %d)", task.Id), nil
	}
	return fmt.Sprintf("task added failed (ID: %d)", task.Id), nil
}

func (ts *taskService) UpdateTask(id int, desc string) string {
	ok, err := ts.TaskRepository.UpdateTask(id, desc, filename)
	helper.LogIfError(ts.Log, err)
	if ok {
		return fmt.Sprintf("Update successfully (ID: %d)", id)
	}
	return fmt.Sprintf("Update failed (ID: %d)", id)
}

func (ts *taskService) DeleteTask(id int) string {
	ok, err := ts.TaskRepository.DeleteTask(id, filename)
	helper.LogIfError(ts.Log, err)
	if ok {
		return fmt.Sprintf("Delete successfully (ID: %d)", id)
	}
	return fmt.Sprintf("Delete failed (ID: %d)", id)
}

func (ts *taskService) MarkTask(status string, id int) string {
	if status == helper.IN_PROGRESS {
		ok, err := ts.TaskRepository.MarkTask(status, id, filename)
		helper.LogIfError(ts.Log, err)
		if ok {
			return fmt.Sprintf("mark in-progress task successfully (ID: %d)", id)
		}
		return fmt.Sprintf("mark in-progress task failed (ID: %d)", id)
	}

	if status == helper.DONE {
		ok, err := ts.TaskRepository.MarkTask(status, id, filename)
		helper.LogIfError(ts.Log, err)
		if ok {
			return fmt.Sprintf("mark done task successfully (ID: %d)", id)
		}
		return fmt.Sprintf("mark done task failed (ID: %d)", id)
	}
	return ""
}

func (ts *taskService) AllTasks() map[string][]entities.Task {
	allTask, err := ts.TaskRepository.GetAllTask(filename)
	helper.LogIfError(ts.Log, err)
	return allTask
}

func (ts *taskService) GetTaskByStatus(status string) map[string][]entities.Task {
	tasks, err := ts.TaskRepository.GetTaskByStatus(status, filename)
	helper.LogIfError(ts.Log, err)
	return tasks
}
