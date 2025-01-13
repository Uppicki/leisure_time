package repositories

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/store"
	"sync"
)

type baseRepository struct {
	config *config.RepositoryConfig

	isBindingStores bool
	mu              *sync.RWMutex
}

func (repos *baseRepository) GetType() config.RepositoryType {
	repos.mu.RLock()
	defer repos.mu.RUnlock()

	return repos.config.Type
}

func (repos *baseRepository) BindStores(store.IStoreProvider) {
	repos.mu.Lock()
	defer repos.mu.Unlock()

	repos.isBindingStores = true
}

func (repos *baseRepository) IsReady() bool {
	repos.mu.RLock()
	defer repos.mu.RUnlock()

	return repos.isBindingStores
}
