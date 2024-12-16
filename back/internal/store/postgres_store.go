package store

import "leisure_time/cmd/config"

type PostgresStore struct {
	config *config.StoreConfig
}

func NewPostgresStore(config *config.StoreConfig) IStore {
	return &PostgresStore{
		config: config,
	}
}
