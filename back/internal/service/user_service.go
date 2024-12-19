package service

import (
	"leisure_time/internal/domain/dto"
	"leisure_time/internal/repositories"
)

type UserService struct {
	baseService
	repo repositories.UserRepository
}

func (service *UserService) GetUserByLogin(login string) (dto.UserDTO, error) {
	return dto.UserDTO{}, nil
}

func (service *UserService) CreateUser(login string, password string) error {
	return nil
}

func (service *UserService) GetUsers() ([]dto.UserDTO, error) {
	return []dto.UserDTO{}, nil
}

func (service *UserService) BindRepositories(
	repoProvider repositories.IRepositoryProvider,
) {
	service.repo = repoProvider.GetUserRepository()

	service.isBindingRepos = true
}

func NewMockUserService() *UserService {
	return &UserService{}
}
