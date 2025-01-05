package storemanager

import (
	"leisure_time/cmd/config"
	"leisure_time/internal/store"
)

type IStoreManager interface {
	store.IStoreProvider
	
	Setup()
	MakeMigrations()
}

func NewStoreManager(cfg *config.StoreManagerConfig) IStoreManager {
	manager := &storeManager{
		config: cfg,
		engines: map[config.SourceType]store.IStoreEngine{},
		stores: map[config.StoreType]store.IStore{},
	}

	return manager
}
 