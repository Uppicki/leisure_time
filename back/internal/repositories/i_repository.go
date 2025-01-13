package repositories

import (
	"leisure_time/cmd/config"
	errs "leisure_time/internal/domain/errors"
	"leisure_time/internal/domain/models"
	"leisure_time/internal/store"
	"sync"
)

type IRepository interface {
	GetType() config.RepositoryType
	BindStores(store.IStoreProvider)
	IsReady() bool
}

type IUserRepository interface {
	GetUserByLogin(string) (models.User, error)
	CreateUser(models.User) error
	GetUsers() ([]models.User, error)
}

type IRepositoryProvider interface {
	GetUserRepository() IUserRepository
}

func RepositoryFactory(cfg *config.RepositoryConfig) (IRepository, error) {
	var repos IRepository

	baseRepo := baseRepository{
		config:          cfg,
		isBindingStores: false,
		mu:              &sync.RWMutex{},
	}

	switch cfg.Type {
	case config.UserRepositoryType:
		repos = &userRepository{
			baseRepository: baseRepo,
		}
	default:
		return nil, errs.UNKNOWN_REPOSITORY_TYPE_ERROR
	}

	return repos, nil
}
