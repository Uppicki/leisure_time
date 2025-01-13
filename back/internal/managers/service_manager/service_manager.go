package servicemanager

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/repositories"
	"leisure_time/internal/service"
	"sync"

	"github.com/sirupsen/logrus"
)

type serviceManager struct {
	config *config.ServiceManagerConfig
	mu     *sync.Mutex

	logger *logrus.Logger

	services map[string]service.IService
}

func (manager *serviceManager) sortConfigs() []config.ServiceConfig {
	//configs := make([]*config.ServiceConfig)
	return manager.config.ServicesConfig
}

func (manager *serviceManager) BindLogger(logger *logrus.Logger) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	manager.logger = logger
}

func (manager *serviceManager) GetUserService() service.IUserService {
	srvc, ok := manager.services[config.UserServiceType.String()]

	if !ok {

	} else if !srvc.IsReady() {

	}

	userService := srvc.(service.IUserService)

	return userService
}

func (manager *serviceManager) SetupAndBinding(
	repoProvider repositories.IRepositoryProvider,
) {
	servicesConfigQueue := manager.sortConfigs()

	for _, serviceConfig := range servicesConfigQueue {
		if _, ok := manager.services[serviceConfig.Type.String()]; !ok {
			service, err := service.ServiceFactory(&serviceConfig)
			if err != nil {
				continue
			}
			service.BindRepositories(repoProvider)
			service.BindServices(manager)

			if service.IsReady() {
				manager.services[service.GetType().String()] = service
			}
		}
	}
}
