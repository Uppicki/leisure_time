package service

import (
	"leisure_time/cmd/config"
	errs "leisure_time/internal/domain/errors"
	"leisure_time/internal/domain/models"
	"leisure_time/internal/repositories"
)

type IService interface {
	GetType() config.ServiceType

	BindRepositories(repositories.IRepositoryProvider)
	BindServices(IServiceProvider)

	IsReady() bool
}

type IUserService interface {
	GetUserByLogin(string) (models.User, error)
	CreateUser(string, string) error
	GetUsers() ([]models.User, error)
}

type IServiceProvider interface {
	GetUserService() IUserService
}

func ServiceFactory(cfg *config.ServiceConfig) (IService, error) {
	switch cfg.Type {
	case config.UserServiceType:
		return &userService{
			baseService: baseService{
				config: cfg,
			},
		}, nil
	default:
		return nil, errs.UNKNOWN_SERVICE_TYPE_ERROR
	}
}
