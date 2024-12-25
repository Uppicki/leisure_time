package storemanager

import (
	config "leisure_time/cmd/config"
	errs "leisure_time/internal/domain/errors"
	stores "leisure_time/internal/store"
)

type storeManager struct {
	config *config.StoreManagerConfig

	engines map[config.SourceType]stores.IStoreEngine
	stores  map[config.StoreType]stores.IStore
}

func (storeManager *storeManager) Setup() {
	for _, cfg := range storeManager.config.StoresConfig {
		eng, engExsits := storeManager.engines[cfg.SourceType]
		if !engExsits {
			if eng = stores.StoreEngineFactory(&cfg); eng != nil {
				storeManager.engines[cfg.SourceType] = eng
			} else {
				continue
			}
		}

		if store := stores.StoreFactory(
			&cfg,
			eng,
		); store != nil {
			storeManager.stores[cfg.StoreType] = store
		}
	}
}

func (storeManager *storeManager) MakeMigrations() {
	for _, store := range storeManager.stores {
		store.BindModels()

	}
}

func (storeManager *storeManager) GetUserStore() (stores.IUserStore, error) {
	store := storeManager.stores[config.USER_TYPE]

	if !store.IsReady() {
		return nil, errs.STORE_UNREADY_ERROR
	}

	return store.(stores.IUserStore), nil
}
