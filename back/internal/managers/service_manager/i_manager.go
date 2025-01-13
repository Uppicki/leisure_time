package servicemanager

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/repositories"
	"leisure_time/internal/service"
	"sync"

	"github.com/sirupsen/logrus"
)

type IServiceManager interface {
	service.IServiceProvider
	BindLogger(logger *logrus.Logger)
	SetupAndBinding(repositories.IRepositoryProvider)
}

func NewServiceManager(cfg *config.ServiceManagerConfig) IServiceManager {
	return &serviceManager{
		config:   cfg,
		mu:       &sync.Mutex{},
		services: make(map[string]service.IService),
	}
}
