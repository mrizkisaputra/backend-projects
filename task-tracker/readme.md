# Task tracker project #
Sample solution for the [Task Tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/). Task tracker is a project used to track and manage your tasks.
The application run from the command line (CLI), accept user actions and input as arguments, and store the tasks in a JSON file. Some of the features in this application include:

- Add, Update, and Delete Task
- Mark a task as (in-progress, done)
- List all Tasks
- List all Tasks that are (done, in-progress, todo)

## How to run the project? #
1. Install **[Go language](https://go.dev/)**. Check installation ```go  version```
2. Download **ZIP** file or use the GIT ```git clone https://github.com/mrizkisaputra/backend-projects.git```
3. Open terminal, navigate to directory project **task-tracker** ```cd task-tracker```
4. Run the app ```go run main.go```

## Example #
List of commands
```shell
# Add, Update, and Delete Task
go run main.go add "learning golang" # Output: task added successfully (ID: 1)
go run main.go update 1 "learning GO" # Output: update successfully (ID: 1)
go run main.go delete 1 # Output: delete successfully (ID: 1)

# Mark a task as in-progress or done (default: todo)
go run main.go mark-in-progress 1
go run main.go mark-done 1

# List all Tasks
go run main.go list

# List all Tasks that are (done, in-progress, todo)
go run main.go list done
go run main.go list in-progress
go run main.go list todo
```