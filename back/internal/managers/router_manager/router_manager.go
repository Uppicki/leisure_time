package routermanager

import (
	"leisure_time/cmd/config"
	servicemanager "leisure_time/internal/managers/service_manager"
	routers "leisure_time/internal/router"
	"sync"

	"github.com/sirupsen/logrus"
)

type routerManager struct {
	config *config.RouterManagerConfig
	mu     *sync.Mutex

	logger *logrus.Logger

	routers map[string]routers.IRouter
}

func (manager *routerManager) SetupAndBinding(
	serviceManager servicemanager.IServiceManager,
) {
	for _, cfg := range manager.config.RoutersConfig {
		name := cfg.Type.String() + "|" + cfg.Port

		if _, ok := manager.routers[name]; ok {
			continue
		}

		if router, err := routers.RouterFactory(&cfg); err == nil {
			router.BindingServiceManager(serviceManager)
			router.Setup()

			manager.routers[name] = router
		}
	}
}

func (manager *routerManager) BindLogger(logger *logrus.Logger) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	manager.logger = logger
}

func (manager *routerManager) Start() {
	for _, router := range manager.routers {
		router.Start()
	}
}
