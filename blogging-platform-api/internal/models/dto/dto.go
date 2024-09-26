package dto

type APIResponse[T any] struct {
	Status int           `json:"status"`
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	TotalData int64 `json:"total_data"`
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
	TotalPage int64 `json:"total_page"`
}

type APIResponseError struct {
	Status int         `json:"status"`
	Errors interface{} `json:"errors"`
}
