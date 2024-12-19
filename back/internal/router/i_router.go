package router

import (
	"leisure_time/cmd/config"
	errs "leisure_time/internal/domain/errors"
	"leisure_time/internal/service"
)

type IRouter interface {
	BindingServiceManager(service.IServiceProvider)
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
