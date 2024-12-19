package servicemanager

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/repositories"
	"leisure_time/internal/service"
)

type serviceManager struct {
	config   *config.ServiceManagerConfig
	services map[string]service.IService
}

func (manager *serviceManager) sortConfigs() []config.ServiceConfig {
	//configs := make([]*config.ServiceConfig)
	return manager.config.ServicesConfig
}

func (manager *serviceManager) GetUserService() *service.UserService {
	return &service.UserService{}
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
