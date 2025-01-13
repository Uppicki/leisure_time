package routermanager

import (
	"leisure_time/cmd/config"
	servicemanager "leisure_time/internal/managers/service_manager"
	routers "leisure_time/internal/router"

	"github.com/sirupsen/logrus"
)

type IRouterManager interface {
	SetupAndBinding(servicemanager.IServiceManager)
	BindLogger(logger *logrus.Logger)
	Start()
}

func NewRouterManager(cfg *config.RouterManagerConfig) IRouterManager {

	routerManager := &routerManager{
		config:  cfg,
		routers: make(map[string]routers.IRouter),
	}

	return routerManager
}
