package dto

type ApiResponse struct {
	Status int           `json:"status"`
	Data   interface{}   `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
}

type ApiResponseError[T interface{}] struct {
	Status    int    `json:"status"`
	Timestamp string `json:"timestamp"`
	Message   string `json:"message"`
	Error     []T    `json:"errors,omitempty"`
}

type PageMetadata struct {
	TotalData int64 `json:"total_data"`
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	TotalPage int64 `json:"total_page"`
}

type ApiValidationError struct {
	Field         string `json:"field"`
	RejectedValue any    `json:"rejected_value"`
	Message       string `json:"message"`
}
