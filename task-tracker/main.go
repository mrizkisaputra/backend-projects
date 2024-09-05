package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"task-tracker/helper"
	"task-tracker/repositories"
	"task-tracker/services"
	"time"
)

var Log = logrus.New()

func setupLogger() {
	Log.SetLevel(logrus.ErrorLevel)
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
}

func main() {
	setupLogger()

	tr := repositories.NewTaskRepository(Log)
	ts := services.NewTaskService(Log, tr)

	/* accept user actions and inputs as arguments from CLI */
	if len(os.Args) < 2 {
		fmt.Println("Please provide command [add,update,delete,list,mark-in-progress]")
		os.Exit(1)
	}

	command := os.Args[1:]
	switch command[0] {
	case "add":
		{
			result, _ := ts.AppendTask(strings.Join(command[1:], " "))
			fmt.Println(result)
		}
	case "update":
		{
			parseInt, err := strconv.Atoi(command[1])
			helper.LogIfError(Log, err)
			result := ts.UpdateTask(parseInt, strings.Join(command[2:], " "))
			fmt.Println(result)
		}
	case "delete":
		{
			parseInt, err := strconv.Atoi(command[1])
			helper.LogIfError(Log, err)
			result := ts.DeleteTask(parseInt)
			fmt.Println(result)
		}
	case "mark-in-progress":
		{
			parseInt, err := strconv.Atoi(command[1])
			helper.LogIfError(Log, err)
			result := ts.MarkTask(helper.IN_PROGRESS, parseInt)
			fmt.Println(result)
		}
	case "mark-done":
		{
			parseInt, err := strconv.Atoi(command[1])
			helper.LogIfError(Log, err)
			result := ts.MarkTask(helper.DONE, parseInt)
			fmt.Println(result)
		}
	case "list":
		{
			if len(command) == 2 {
				if command[1] == helper.DONE {
					tasks := ts.GetTaskByStatus(helper.DONE)
					fmt.Println(tasks)
					return
				}

				if command[1] == helper.TODO {
					tasks := ts.GetTaskByStatus(helper.TODO)
					fmt.Println(tasks)
					return
				}

				if command[1] == helper.IN_PROGRESS {
					tasks := ts.GetTaskByStatus(helper.IN_PROGRESS)
					fmt.Println(tasks)
					return
				}
			}
			tasks := ts.AllTasks()
			fmt.Println(tasks)
		}
	default:
		{
			fmt.Printf("command %s not valid, please enter the correct command", command[0])
		}
	}
}
