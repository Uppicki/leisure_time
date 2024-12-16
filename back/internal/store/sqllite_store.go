package store

import "leisure_time/cmd/config"

type SQLLiteStore struct {
	config *config.StoreConfig
}

func NewSQLLiteStore(config *config.StoreConfig) IStore {
	return &SQLLiteStore{
		config: config,
	}
}
