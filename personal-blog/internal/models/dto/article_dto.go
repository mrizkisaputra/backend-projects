package dto

type ArticlePayloadBodyRequest struct {
	Title       string `json:"title" validate:"required"`
	Content     string `json:"content" validate:"required"`
	PublishDate string `json:"publish_date"`
	Category    string `json:"category" validate:"alpha"`
	Tags        string `json:"tags"`
}

type ArticlePayloadParamIdRequest struct {
	Id int64 `validate:"numeric"`
}
