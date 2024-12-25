package repositorymanager

import "leisure_time/internal/repositories"

type IRepositoryManager interface {
	repositories.IRepositoryProvider
	Setup()
	Run()
}
