package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	. "task-tracker/entities"
	"task-tracker/helper"
	"time"
)

type taskRepository struct {
	Log *logrus.Logger
}

func NewTaskRepository(log *logrus.Logger) TaskRepositoryInterface {
	return &taskRepository{
		Log: log,
	}
}

func (tr *taskRepository) CreateFile(filename string) (string, error) {
	/* detect whether the file already exists */
	_, err := os.Stat(filename)

	/* create a new file, if it does not already exist */
	if os.IsNotExist(err) {

		file, errAbs := filepath.Abs(filename)
		if errAbs != nil {
			return "", errAbs
		}

		/* create file */
		if file, err := os.Create(file); err != nil {
			tr.Log.Errorf("error creating file: %v", err.Error())
			return "", fmt.Errorf("error while creating file because: %v", err.Error())
		} else {
			defer func() {
				err := file.Close()
				if err != nil {
					tr.Log.Errorf("error closing file: %v", err.Error())
				}
			}()
			tr.WriteFile(filename, map[string][]Task{"tasks": {}})
			tr.Log.Tracef("created file: %s", file.Name())
			return filename, nil
		}
	}
	tr.Log.Trace("filename already exists")
	return "", nil
}

func (tr *taskRepository) WriteFile(filename string, tasks map[string][]Task) error {
	fileLocation, errAbs := filepath.Abs(filename)
	if errAbs != nil {
		return errAbs
	}

	/* open file with access level READ & WRITE */
	file, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		tr.Log.Errorf("error opening file: %v", err.Error())
		return fmt.Errorf("error while opening file because: %v", err.Error())
	}

	errEncode := json.NewEncoder(file).Encode(&tasks)
	if errEncode != nil {
		tr.Log.Errorf("error encode tasks to json: %v", errEncode.Error())
		return fmt.Errorf("error encode task to json: %v", errEncode.Error())
	}
	return nil
}

func (tr *taskRepository) ReadFile(filename string) (map[string][]Task, error) {
	fileLocation, errAbs := filepath.Abs(filename)
	if errAbs != nil {
		return nil, errAbs
	}
	file, err := os.OpenFile(fileLocation, os.O_RDONLY, 0644)
	if err != nil {
		tr.Log.Errorf("error opening file: %v", err.Error())
		return nil, err
	}

	var tasks map[string][]Task
	errDecode := json.NewDecoder(file).Decode(&tasks)
	if errDecode != nil {
		tr.Log.Errorf("error decoding file: %v", errDecode.Error())
		return nil, errDecode
	}
	return tasks, nil
}

func (tr *taskRepository) AppendTask(task Task, filename string) (bool, error) {
	readTask, errReadFile := tr.ReadFile(filename)
	if errReadFile != nil {
		tr.Log.Errorf("error reading file: %v", errReadFile.Error())
		return false, errReadFile
	}

	readTask["tasks"] = append(readTask["tasks"], task)
	errWriteFile := tr.WriteFile(filename, readTask)
	if errWriteFile != nil {
		tr.Log.Errorf("error append task to file: %v", errWriteFile.Error())
		return false, errWriteFile
	}
	return true, nil
}

func (tr *taskRepository) UpdateTask(id int, desc string, filename string) (bool, error) {
	task, err := tr.ReadFile(filename)
	if err != nil {
		tr.Log.Errorf("error reading file to update: %v", err.Error())
		return false, nil
	}

	/* search task by id */
	for i, e := range task["tasks"] {
		if e.Id == id {
			parsedTime, err := time.Parse(time.RFC3339, time.Now().Local().Format(time.RFC3339))
			if err != nil {
				return false, err
			}
			newTask := Task{
				Id:          task["tasks"][i].Id,
				Description: desc,
				Status:      task["tasks"][i].Status,
				CreatedAt:   task["tasks"][i].CreatedAt,
				UpdatedAt:   parsedTime,
			}

			/* update task */
			task["tasks"][i] = newTask
			if err := tr.WriteFile(filename, task); err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func (tr *taskRepository) DeleteTask(id int, filename string) (bool, error) {
	task, err := tr.ReadFile(filename)
	if err != nil {
		return false, err
	}

	for i, e := range task["tasks"] {
		if e.Id == id {
			task["tasks"][i] = Task{}
			if err := tr.WriteFile(filename, task); err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func (tr *taskRepository) MarkTask(status string, id int, filename string) (bool, error) {
	task, err := tr.ReadFile(filename)
	if err != nil {
		return false, err
	}

	for i, e := range task["tasks"] {
		if e.Id == id {
			newTask := Task{
				Id:          task["tasks"][i].Id,
				Description: task["tasks"][i].Description,
				Status:      status,
				CreatedAt:   task["tasks"][i].CreatedAt,
				UpdatedAt:   task["tasks"][i].UpdatedAt,
			}
			task["tasks"][i] = newTask

			if err := tr.WriteFile(filename, task); err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}

func (tr *taskRepository) GetAllTask(filename string) (map[string][]Task, error) {
	tasks, err := tr.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *taskRepository) GetTaskByStatus(status string, filename string) (map[string][]Task, error) {
	tasks, err := tr.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasksByStatus = map[string][]Task{"tasks": {}}
	for i, e := range tasks["tasks"] {
		if status == helper.IN_PROGRESS && status == e.Status {
			tasksByStatus["tasks"] = append(tasksByStatus["tasks"], tasks["tasks"][i])
			continue
		}

		if status == helper.DONE && status == e.Status {
			tasksByStatus["tasks"] = append(tasksByStatus["tasks"], tasks["tasks"][i])
			continue
		}

		if status == helper.TODO && status == e.Status {
			tasksByStatus["tasks"] = append(tasksByStatus["tasks"], tasks["tasks"][i])
			continue
		}
	}

	return tasksByStatus, nil
}
