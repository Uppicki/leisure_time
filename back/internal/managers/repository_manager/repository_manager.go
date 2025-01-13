package repositorymanager

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/repositories"
	"leisure_time/internal/store"
	"sync"

	"github.com/sirupsen/logrus"
)

type repositoryManager struct {
	config *config.RepositoryManagerConfig
	mu     *sync.RWMutex

	logger *logrus.Logger

	repositories map[string]repositories.IRepository
}

func (manager *repositoryManager) BindLogger(logger *logrus.Logger) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	manager.logger = logger
}

func (manager *repositoryManager) sortConfigs() []config.RepositoryConfig {
	manager.mu.Lock()
	manager.mu.Unlock()

	return manager.config.ReposConfig
}

func (manager *repositoryManager) SetupAndBinding(
	provider store.IStoreProvider,
) {
	reposConfigQueue := manager.sortConfigs()

	manager.mu.Lock()
	defer manager.mu.Unlock()

	for _, reposConfig := range reposConfigQueue {
		if _, ok := manager.repositories[reposConfig.Type.String()]; !ok {
			repos, err := repositories.RepositoryFactory(&reposConfig)
			if err != nil {
				continue
			}

			repos.BindStores(provider)

			if repos.IsReady() {
				manager.repositories[reposConfig.Type.String()] = repos
			}
		}
	}

}

func (manager *repositoryManager) GetUserRepository() repositories.IUserRepository {
	manager.mu.RLock()
	defer manager.mu.RUnlock()

	repo, ok := manager.repositories[config.UserRepositoryType.String()]
	if !ok {
		return nil
	}

	userRepo := repo.(repositories.IUserRepository)

	return userRepo
}
