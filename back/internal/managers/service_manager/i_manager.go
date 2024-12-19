package servicemanager

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/repositories"
	"leisure_time/internal/service"
)

type IServiceManager interface {
	service.IServiceProvider
	SetupAndBinding(repositories.IRepositoryProvider)
}

func NewServiceManager(cfg *config.ServiceManagerConfig) IServiceManager {
	return nil
}
