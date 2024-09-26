package converters

import (
	"blogging-platform-api/internal/models/dto"
	"blogging-platform-api/internal/models/entities"
	"strings"
)

func PostToPostResponse(entity *entities.Post) *dto.PostResponse {
	return &dto.PostResponse{
		Id:        entity.Id,
		Title:     entity.Title,
		Content:   entity.Content,
		Category:  entity.Category,
		Tags:      strings.Split(entity.Tags, ","),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
