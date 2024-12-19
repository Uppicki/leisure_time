package service

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/repositories"
)

type baseService struct {
	config *config.ServiceConfig

	isBindingRepos    bool
	isBindingServices bool
}

func (service *baseService) GetType() config.ServiceType {
	return service.config.Type
}

func (service *baseService) BindRepositories(repositories.IRepositoryProvider) {
	service.isBindingRepos = true
}

func (service *baseService) BindServices(IServiceProvider) {
	service.isBindingServices = true
}

func (service *baseService) IsReady() bool {
	return service.isBindingRepos && service.isBindingServices
}
