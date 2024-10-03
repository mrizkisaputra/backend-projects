package dto

type CreateTaskRequest struct {
	IdUser      string `json:"-" validate:"required,max=100"`
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=200"`
}

type UpdateTaskRequest struct {
	IdUser      string `json:"-" validate:"required,max=100"`
	Id          string `json:"-" validate:"required,max=100"`
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"required,max=200"`
}

type UpdateStatusTaskRequest struct {
	IdUser string `json:"-" validate:"required,max=100"`
	Id     string `json:"-" validate:"required,max=100"`
	Mark   string `json:"-" validate:"max=12,oneof=in-progress done"`
}

type DeleteTaskRequest struct {
	IdUser string `json:"-" validate:"required,max=100"`
	Id     string `json:"id" validate:"required,max=100"`
}

type SearchTaskRequest struct {
	IdUser string `validate:"required,max=100"`
	Status string `validate:"max=12,oneof=in-progress done"`
	Sort   string `validate:"oneof=created_at"`
	Order  string `validate:"oneof=asc desc"`
	Page   int    `validate:"min=0"`
	Limit  int    `validate:"max=100"`
}

type TaskResponse struct {
	Id          string `json:"id"`
	IdUser      string `json:"id_user"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
