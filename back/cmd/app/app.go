package app

import (
	"leisure_time/cmd/config"
	repositorymanager "leisure_time/internal/managers/repository_manager"
	routermanager "leisure_time/internal/managers/router_manager"
	servicemanager "leisure_time/internal/managers/service_manager"
	storemanager "leisure_time/internal/managers/store_manager"
)

type MyApp struct {
	config *config.AppConfig

	routerManager     routermanager.IRouterManager
	serviceManager    servicemanager.IServiceManager
	repositoryManager repositorymanager.IRepositoryManager
	storeManager      storemanager.IStoreManager
}

func (app *MyApp) Setup() {
}

func (app *MyApp) Run() {

}

func NewApp(config *config.AppConfig) *MyApp {
	return &MyApp{
		config: config,
	}
}
