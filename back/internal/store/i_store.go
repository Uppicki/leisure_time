package store

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/domain/models"

	"gorm.io/gorm"
)

type IStore interface {
	BindModels() error
	bindModel(any) error

	IsReady() bool

	DB() (*gorm.DB, error)
}

type IStoreEngine interface {
	BindModel(any) error
	Migrate() error

	Connect() error

	IsConnected() bool
	IsReady() bool

	DB() *gorm.DB
}

type IUserStore interface {
	IStore

	CreateUser(models.User) error
	GetUserByLogin(string) (models.User, error)
	GetUserList() ([]models.User, error)
}

type IStoreProvider interface {
	GetUserStore() IUserStore
}

func StoreFactory(cfg *config.StoreConfig) IStore {
	return nil
}
