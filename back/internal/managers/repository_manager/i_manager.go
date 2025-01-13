package repositorymanager

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/repositories"
	"leisure_time/internal/store"
	"sync"
)

type IRepositoryManager interface {
	repositories.IRepositoryProvider
	SetupAndBinding(store.IStoreProvider)
}

func NewRepositoryManager(
	cfg *config.RepositoryManagerConfig,
) IRepositoryManager {
	manager := &repositoryManager{
		config:       cfg,
		mu:           &sync.RWMutex{},
		repositories: make(map[string]repositories.IRepository),
	}

	return manager
}
