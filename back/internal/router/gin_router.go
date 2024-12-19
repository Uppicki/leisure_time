package router

import (
	config "leisure_time/cmd/config"
	ginhandlers "leisure_time/internal/gin_handlers"
	"leisure_time/internal/service"

	gin "github.com/gin-gonic/gin"
)

type ginRouter struct {
	config          *config.RouterConfig
	engine          *gin.Engine
	serviceProvider service.IServiceProvider

	isReady  bool
	isActive bool
}

func (router *ginRouter) setupRoutes() {
	if router.isBindingService() {
		baseGroup := router.engine.Group("")

		userService := router.serviceProvider.GetUserService()

		userHandler := ginhandlers.NewUserHandler(userService)

		usersGroup := baseGroup.Group("/users")
		{
			usersGroup.POST("", userHandler.CreateUser)
			usersGroup.GET("", userHandler.GetUsers)
			usersGroup.GET("/:login", userHandler.GetUserByLogin)
		}
	}
}

func (router *ginRouter) setReady() {
	if router.isBindingService() && router.engine != nil {
		router.isReady = true
	}
}

func (router *ginRouter) isBindingService() bool {
	return router.serviceProvider != nil
}

func (router *ginRouter) BindingServiceManager(
	serviceProvider service.IServiceProvider,
) {
	router.serviceProvider = serviceProvider
}

func (router *ginRouter) Setup() {
	router.setupRoutes()

	router.setReady()
}

func (router *ginRouter) Start() {
	if router.isReady {
		router.isActive = true
		go router.engine.Run(":" + router.config.Port)
	}
}
