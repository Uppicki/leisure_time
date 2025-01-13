package app

import (
	"leisure_time/cmd/config"
	repositorymanager "leisure_time/internal/managers/repository_manager"
	routermanager "leisure_time/internal/managers/router_manager"
	servicemanager "leisure_time/internal/managers/service_manager"
	storemanager "leisure_time/internal/managers/store_manager"

	"github.com/sirupsen/logrus"
)

type MyApp struct {
	config *config.AppConfig

	logger *logrus.Logger

	routerManager     routermanager.IRouterManager
	serviceManager    servicemanager.IServiceManager
	repositoryManager repositorymanager.IRepositoryManager
	storeManager      storemanager.IStoreManager
}

func (app *MyApp) Setup() {
	app.storeManager.Setup()
	app.storeManager.MakeMigrations()
	app.repositoryManager.SetupAndBinding(app.storeManager)
	app.serviceManager.SetupAndBinding(app.repositoryManager)
	app.routerManager.SetupAndBinding(app.serviceManager)
}

func (app *MyApp) Start() {
	app.routerManager.Start()
}

func NewApp(cfg *config.AppConfig) *MyApp {
	logger := logrus.New()

	switch cfg.Mod {
	case config.Debug:
		logger.SetLevel(logrus.DebugLevel)
	case config.Prod:
		logger.SetLevel(logrus.InfoLevel)
	}

	routerManager := routermanager.NewRouterManager(&cfg.RouterManagerConfig)
	serviceManager := servicemanager.NewServiceManager(&cfg.ServiceManagerConfig)
	repoManager := repositorymanager.NewRepositoryManager(&cfg.RepositoryManagerConfig)
	storeManager := storemanager.NewStoreManager(&cfg.StoreManagerConfig)

	app := &MyApp{
		config:            cfg,
		logger:            logger,
		routerManager:     routerManager,
		serviceManager:    serviceManager,
		repositoryManager: repoManager,
		storeManager:      storeManager,
	}

	app.storeManager.BindLogger(app.logger)
	app.repositoryManager.BindLogger(app.logger)
	app.serviceManager.BindLogger(app.logger)
	app.routerManager.BindLogger(app.logger)

	return app
}
