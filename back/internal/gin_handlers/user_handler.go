package ginhandlers

import (
	service "leisure_time/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.IUserService
}

func (handler *UserHandler) GetUserByLogin(ctx *gin.Context) {
	login := ctx.Param("login")

	user, err := handler.userService.GetUserByLogin(login)

	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"err": "Bad",
			},
		)
	}

	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"mes":  "OK",
			"user": user,
		},
	)
}

func (handler *UserHandler) CreateUser(ctx *gin.Context) {
	var request createUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		return
	}

	err := handler.userService.CreateUser(
		request.Login,
		request.Password,
	)

	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"err": "Bad",
			},
		)
	}

	ctx.IndentedJSON(
		http.StatusOK,
		gin.H{
			"mes": "OK",
		},
	)
}

func (handler *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := handler.userService.GetUsers()

	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			gin.H{
				"err": "Bad",
			},
		)
	}

	ctx.IndentedJSON(
		http.StatusOK,
		userListResponse{
			Users: users,
		},
	)
}

func NewUserHandler(service service.IUserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}
