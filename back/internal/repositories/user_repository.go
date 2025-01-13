package repositories

import (
	"leisure_time/internal/domain/models"
	"leisure_time/internal/store"
)

type userRepository struct {
	baseRepository

	userStore store.IUserStore
}

func (repos *userRepository) BindStores(storeProvider store.IStoreProvider) {
	userStore, err := storeProvider.GetUserStore()

	if err != nil {
		return
	}

	repos.mu.Lock()
	defer repos.mu.Unlock()

	repos.userStore = userStore
}

func (repos *userRepository) GetUserByLogin(login string) (models.User, error) {
	return repos.userStore.GetUserByLogin(login)
}

func (repos *userRepository) CreateUser(user models.User) error {
	return repos.userStore.CreateUser(user)
}

func (repos *userRepository) GetUsers() ([]models.User, error) {
	return repos.userStore.GetUserList()
}
