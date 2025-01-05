package store

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/domain/models"

	"gorm.io/driver/sqlite"
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
	GetUserStore() (IUserStore, error)
}

func StoreFactory(cfg *config.StoreConfig, eng IStoreEngine) IStore {
	store := &store{
		eng: eng,
	}

	switch cfg.StoreType{
	case config.USER_TYPE:
		return &userStore{
			IStore: store,
		}
	default:
		return nil
	}
}

func StoreEngineFactory(cfg *config.StoreConfig) IStoreEngine {
	var dialector gorm.Dialector

	switch cfg.SourceDialect{
	case config.SQLLITE_DIALECT:
		dialector = sqlite.Open("gorm.db")
	default:
		return nil
	}

	switch cfg.SourceType {
	case config.GIN_SOURCE:
		engine := &gormEngine{
			dialector: dialector,
		}
		
		return engine
	default:
		return nil
	}
}
