package mockhandler

import (
	repositoriesmanager "leisure_time_back/cmd/repositories_manager"
	"leisure_time_back/internal/payload/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MockHandler struct {
}

func NewMockHandler() *MockHandler {
	handler := MockHandler{}

	return &handler
}

func (handler *MockHandler) Hello(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		responses.HelloResponse{
			Message: "Hello World",
		},
	)
}

func (handler *MockHandler) BindingRepository(
	repo repositoriesmanager.IRepository,
) {
	return
}
