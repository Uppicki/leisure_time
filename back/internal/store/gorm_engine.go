package store

import (
	errs "leisure_time/internal/domain/errors"
	"sync"

	"gorm.io/gorm"
)

type gormEngine struct {
	dialector gorm.Dialector

	db     *gorm.DB
	models map[any]bool

	mu      sync.Mutex
	isReady bool
}

func (eng *gormEngine) Connect() error {
	db, err := gorm.Open(eng.dialector, &gorm.Config{})

	if err != nil {
		return err
	}

	eng.mu.Lock()
	defer eng.mu.Unlock()

	if eng.db != nil {
		return errs.STORE_CONNECTION_IS_DEFINED_ERROR
	}

	eng.db = db

	return nil
}

func (eng *gormEngine) BindModel(model any) error {
	eng.mu.Lock()
	defer eng.mu.Unlock()

	if _, isExsist := eng.models[model]; isExsist {
		return errs.STORE_MODEL_EXSITS_ERROR
	}

	eng.models[model] = true
	return nil
}

func (eng *gormEngine) Migrate() error {
	if !eng.IsConnected() {
		return errs.STORE_CONNECTION_UNDEFINED_ERROR
	}

	eng.mu.Lock()
	defer eng.mu.Unlock()

	modelList := make([]any, len(eng.models))

	for model := range eng.models {
		modelList = append(modelList, model)
	}

	if err := eng.db.AutoMigrate(modelList...); err != nil {
		return errs.STORE_MIGRATION_ERROR
	}

	eng.isReady = true
	return nil
}

func (eng *gormEngine) IsConnected() bool {
	eng.mu.Lock()
	defer eng.mu.Unlock()
	return eng.db != nil
}

func (eng *gormEngine) IsReady() bool {
	return eng.isReady
}

func (eng *gormEngine) DB() *gorm.DB {
	eng.mu.Lock()
	defer eng.mu.Unlock()

	return eng.db
}
