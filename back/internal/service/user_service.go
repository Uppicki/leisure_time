package service

import (
	"leisure_time/internal/domain/models"
	"leisure_time/internal/repositories"
)

type userService struct {
	baseService
	repo repositories.IUserRepository
}

func (service *userService) GetUserByLogin(login string) (models.User, error) {
	return service.repo.GetUserByLogin(login)
}

func (service *userService) CreateUser(login string, password string) error {
	user := models.User{
		Login:    login,
		Password: password,
	}

	return service.repo.CreateUser(user)
}

func (service *userService) GetUsers() ([]models.User, error) {
	return service.repo.GetUsers()
}

func (service *userService) BindRepositories(
	repoProvider repositories.IRepositoryProvider,
) {
	service.repo = repoProvider.GetUserRepository()

	service.isBindingRepos = true
}
