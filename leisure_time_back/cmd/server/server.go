package server

import (
	handlersmanager "leisure_time_back/cmd/handlers_manager"
	repositoriesmanager "leisure_time_back/cmd/repositories_manager"
	"leisure_time_back/cmd/router"
)

type Server struct {
	handlersManager     *handlersmanager.HandlersManager
	repositoriesManager *repositoriesmanager.RepositoryManager
	router              *router.Router
}

func NewServer() *Server {
	router := router.NewRouter()
	repositoriesManager := repositoriesmanager.NewRepositoryManager()
	handlersManager := handlersmanager.NewHandlersManager()

	server := Server{
		handlersManager:     handlersManager,
		repositoriesManager: repositoriesManager,
		router:              router,
	}

	return &server
}

func (server *Server) SetupServer() {
	//server.repositoriesManager.SetupRepositories()
	server.handlersManager.SetupHandlers()
	server.router.SetupRoutes(server.handlersManager)
}

func (server *Server) RunServer() {
	server.router.Run()
}
