# Todo List API #
Sample solution for challenge [todolist api](https://roadmap.sh/projects/todo-list-api) from [roadmap.sh](https://roadmap.sh)

## Features
1. User registration to create a new user
2. Login endpoint to authenticate the user and generate a token
3. CRUD operations for managing the to-do list
4. Implement pagination, filtering, and sorting for the to-do list

## Tech Stack
-  [Golang](https://go.dev/) (programming language)
-  [MySQL](https://www.mysql.com/downloads/) (relational database)
-  [Docker](https://www.docker.com/) (container)

## Framework & Library
-  [GORM](https://gorm.io/docs/index.html) (ORM)
-  [Logrus](https://github.com/sirupsen/logrus) (logger)
-  [Viper](https://github.com/spf13/viper) (configuration)
-  [GoFiber](https://docs.gofiber.io/) (http framework)
-  [Go Playground Validator](https://github.com/go-playground/validator) (validation)
-  [Golang Migrate](https://github.com/golang-migrate/migrate) (database migration)
-  [Testify](https://github.com/stretchr/testify) (testing)


## Clean Architecture
The project uses Clean Architecture by Robert C. Martin (Uncle Bob).
### Layer of Clean Architecture
-  models
-  respository
-  usecase
-  delivery

## API Documentation
API spec is in [docs](./docs) directory

## Run Application
