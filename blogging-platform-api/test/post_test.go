package test

import (
	"blogging-platform-api/internal/models/dto"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPost_Create_Success(t *testing.T) {
	requestBody := dto.CreatePostRequestBody{
		Title:    "Tutorial Golang",
		Content:  "golang is a programming language",
		Category: "programming",
		Tags:     []string{"Tech", "tutorial"},
	}
	byteBody, errEncoded := json.Marshal(&requestBody)
	assert.Nil(t, errEncoded)

	request := httptest.NewRequest(http.MethodPost, "/api/posts", strings.NewReader(string(byteBody)))
	request.Header.Add("Content-Type", "application/json")
	response, errHttpResponse := app.Test(request)
	assert.Nil(t, errHttpResponse)

	responseBody, errResponseBody := io.ReadAll(response.Body)
	assert.Nil(t, errResponseBody)

	apiResponse := new(dto.APIResponse[dto.PostResponse])
	errDecoded := json.Unmarshal(responseBody, apiResponse)
	assert.Nil(t, errDecoded)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, http.StatusCreated, apiResponse.Status)
	assert.NotNil(t, apiResponse.Data.Id)
	assert.NotNil(t, apiResponse.Data.CreatedAt)
	assert.NotNil(t, apiResponse.Data.UpdatedAt)
	assert.Equal(t, requestBody.Title, apiResponse.Data.Title)
	assert.Equal(t, requestBody.Content, apiResponse.Data.Content)
	assert.Equal(t, requestBody.Category, apiResponse.Data.Category)
	assert.Equal(t, requestBody.Tags, apiResponse.Data.Tags)
}

func TestPost_GetAll(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/posts", nil)
	request.Header.Add("Content-Type", "application/json")

	response, err1 := app.Test(request)
	assert.Nil(t, err1)
	bytes, err2 := io.ReadAll(response.Body)
	assert.Nil(t, err2)

	var apiResponse dto.APIResponse[[]dto.PostResponse]
	err3 := json.Unmarshal(bytes, &apiResponse)
	assert.Nil(t, err3)

	assert.Equal(t, http.StatusOK, apiResponse.Status)
}

func TestPost_Delete(t *testing.T) {
	//requestParam := new(dto.DeletePostRequestParam)
	//db.Where("id = ?").Take(&requestParam)
	//request := httptest.NewRequest(http.MethodDelete, "/api/posts/"+requestParam.Id, nil)
}

func TestPost_Update(t *testing.T) {

}

func TestPost_FindAll(t *testing.T) {

}
