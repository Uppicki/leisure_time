package handlersmanager

import (
	mockhandler "leisure_time_back/internal/handlers/mock_handler"
)

type HandlersManager struct {
	// Types of IHandler
	// 1. MockHandler
	handlers map[EnumHandler]IHandler
}

func NewHandlersManager() *HandlersManager {
	handlersmanager := HandlersManager{
		handlers: make(map[EnumHandler]IHandler),
	}

	handlersmanager.createHandlers()

	return &handlersmanager
}

func (manager *HandlersManager) createHandlers() {
	manager.handlers[MockHandler] = mockhandler.NewMockHandler()
}

func (manager *HandlersManager) SetupHandlers() {
	manager.handlers[MockHandler].BindingRepository(nil)
}

func (manager *HandlersManager) GetHandler(handler EnumHandler) IHandler {
	return manager.handlers[handler]
}
