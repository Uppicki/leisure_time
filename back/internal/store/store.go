package store

import (
	errs "leisure_time/internal/domain/errors"

	"gorm.io/gorm"
)

type store struct {
	eng IStoreEngine
}

func (store *store) BindModels() error {
	return nil
}

func (store *store) bindModel(model any) error {
	return store.eng.BindModel(model)
}

func (store *store) IsReady() bool {
	return store.eng.IsReady()
}

func (store *store) DB() (*gorm.DB, error) {
	if store.IsReady() {
		return store.eng.DB(), nil
	}

	return nil, errs.STORE_UNREADY_ERROR
}
