package repositories

import "leisure_time/internal/domain/models"

type IRepository interface {
	Any()
}

type IUserRepository interface {
	GetUserByLogin(string) (models.User, error)
	CreateUser(models.User) error
	GetUsers() ([]models.User, error)
}

type IRepositoryProvider interface {
	GetUserRepository() IUserRepository
}
