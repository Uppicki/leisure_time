package storemanager

import (
	config "leisure_time/cmd/config"
	stores "leisure_time/internal/store"
)

type StoreManager struct {
	config *config.StoreManagerConfig
	stores []stores.IStore
}

func (storeManager *StoreManager) Setup() {

}
