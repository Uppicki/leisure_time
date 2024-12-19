package service

import (
	"leisure_time/cmd/config"
	errs "leisure_time/internal/domain/errors"
	"leisure_time/internal/repositories"
)

type IService interface {
	GetType() config.ServiceType

	BindRepositories(repositories.IRepositoryProvider)
	BindServices(IServiceProvider)

	IsReady() bool
}

type IServiceProvider interface {
	GetUserService() *UserService
}

func ServiceFactory(cfg *config.ServiceConfig) (IService, error) {
	switch cfg.Type {
	case config.UserServiceType:
		return &UserService{
			baseService: baseService{
				config: cfg,
			},
		}, nil
	default:
		return nil, errs.UNKNOWN_SERVICE_TYPE_ERROR
	}
}
