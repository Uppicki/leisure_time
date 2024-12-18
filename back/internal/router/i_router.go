package router

import (
	"leisure_time/cmd/config"
	errs "leisure_time/internal/domain/errors"
	servicemanager "leisure_time/internal/managers/service_manager"
)

type IRouter interface {
	BindingServiceManager(servicemanager.IServiceManager)
	Setup()
	Start()
}

func RouterFactory(cfg *config.RouterConfig) (IRouter, error) {
	switch cfg.Type {
	case config.GIN:
		return &ginRouter{
			config: cfg,
		}, nil
	case config.GRPC:
		panic("Unimplement yet")
	default:
		return nil, errs.UNKNOWN_ROUTER_TYPE_ERROR
	}
}
