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
	app.storeManager.Setup()
	app.storeManager.MakeMigrations()
	app.serviceManager.SetupAndBinding(app.repositoryManager)
	app.routerManager.SetupAndBinding(app.serviceManager)
}

func (app *MyApp) Start() {
	app.routerManager.Start()
}

func NewApp(cfg *config.AppConfig) *MyApp {
	routerManager := routermanager.NewRouterManager(&cfg.RouterManagerConfig)
	serviceManager := servicemanager.NewServiceManager(&cfg.ServiceManagerConfig)

	storeManager := storemanager.NewStoreManager(&cfg.StoreManagerConfig)

	app := &MyApp{
		config:         cfg,
		routerManager:  routerManager,
		serviceManager: serviceManager,
		storeManager: storeManager,
	}

	return app
}
