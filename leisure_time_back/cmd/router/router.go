package router

import (
	handlersmanager "leisure_time_back/cmd/handlers_manager"
	mockhandler "leisure_time_back/internal/handlers/mock_handler"

	"github.com/gin-gonic/gin"
)

type Router struct {
	ginRouter *gin.Engine
}

func NewRouter() *Router {
	ginRouter := gin.Default()

	router := Router{
		ginRouter: ginRouter,
	}

	return &router
}

func (r *Router) SetupRoutes(handlers *handlersmanager.HandlersManager) {
	r.setupApiV1Routes(handlers)
}

func (r *Router) setupApiV1Routes(handlers *handlersmanager.HandlersManager) {
	group := r.ginRouter.Group("/api/v1")
	{
		group.GET("/hello", handlers.GetHandler(
			handlersmanager.MockHandler,
		).(*mockhandler.MockHandler).Hello,
		)
	}

}

func (r *Router) Run() {
	r.ginRouter.Run(":11111")
}
