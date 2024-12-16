package store

import config "leisure_time/cmd/config"

type LocalStore struct {
	config *config.StoreConfig
}

func NewLocalStore(config *config.StoreConfig) IStore {
	return &LocalStore{
		config: config,
	}
}
