package test

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"task-tracker/entities"
	"task-tracker/helper"
	"task-tracker/repositories"
	"testing"
	"time"
)

var log = logrus.New()

func TestCreateFile(t *testing.T) {
	log.SetLevel(logrus.TraceLevel)
	var ct = repositories.NewTaskRepository(log)
	filename := "tasks.json"
	file, err := ct.CreateFile(filename)
	assert.Nil(t, err)
	assert.Equal(t, filename, file)
}

func TestWriteFile(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	err := ct.WriteFile(
		"tasks.json",
		map[string][]entities.Task{
			"tasks": {
				{Id: 1, Description: "Testing-1", Status: "todo", CreatedAt: time.Now().Local(), UpdatedAt: time.Now()},
			},
		})
	assert.Nil(t, err)
}

func TestReadFile(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	readFile, err := ct.ReadFile("tasks.json")
	assert.Nil(t, err)
	assert.NotNil(t, readFile)
}

func TestAppendFile(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	ok, err := ct.AppendTask(entities.Task{Id: 2, Description: "Testing-2", Status: "in-progress", CreatedAt: time.Now()}, "tasks.json")
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestUpdateTask(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	task, err := ct.UpdateTask(2, "description updated", "tasks.json")

	assert.Nil(t, err)
	assert.True(t, task)
}

func TestDeleteTask(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	task, err := ct.DeleteTask(2, "tasks.json")
	assert.Nil(t, err)
	assert.True(t, task)
}

func TestMarkInProgressTask(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	task, err := ct.MarkTask(helper.IN_PROGRESS, 2, "tasks.json")
	assert.Nil(t, err)
	assert.True(t, task)
}

func TestMarkDoneTask(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	task, err := ct.MarkTask(helper.DONE, 2, "tasks.json")
	assert.Nil(t, err)
	assert.True(t, task)
}

func TestGetAllTask(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	tasks, err := ct.GetAllTask("tasks.json")
	assert.Nil(t, err)
	assert.NotNil(t, tasks)
}

func TestGetTaskByStatus(t *testing.T) {
	var ct = repositories.NewTaskRepository(logrus.New())
	tasks, err := ct.GetTaskByStatus(helper.TODO, "tasks.json")
	assert.Nil(t, err)
	assert.NotNil(t, tasks)
	fmt.Println(tasks)
}
