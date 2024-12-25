package store

import "leisure_time/internal/domain/models"

type userStore struct {
	IStore
}

func (store *userStore) BindModels() error {
	user := &models.User{}

	store.bindModel(user)

	return nil
}

func (store *userStore) CreateUser(user models.User) error {
	db, err := store.DB()
	if err != nil {
		return err
	}

	res := db.Create(&user)

	return res.Error
}

func (store *userStore) GetUsers() (users []models.User, err error) {
	db, err := store.DB()
	if err != nil {
		return users, err
	}

	res := db.Find(&users)

	return users, res.Error
}

func (store *userStore) GetUserByLogin(login string) (models.User, error) {
	var user models.User

	db, err := store.DB()
	if err != nil {
		return user, err
	}

	result := db.Where("login = ?", login).First(&user)

	return user, result.Error
}
