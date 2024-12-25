package config

import appstrings "leisure_time/internal/domain/strings"

type StoreType int

const (
	USER_TYPE StoreType = iota
)

func (storeType StoreType) String() string {
	switch storeType {
	case USER_TYPE:
		return appstrings.USER_TYPE
	default:
		return appstrings.UNKNOWN
	}
}

type SourceType int

const (
	GIN_SOURCE SourceType = iota
)

func (source SourceType) String() string {
	switch source {
	case GIN_SOURCE:
		return appstrings.GIN_SOURCE
	default:
		return appstrings.UNKNOWN_SOURCE
	}
}

type SourceDialect int

const (
	SQLLITE_DIALECT SourceDialect = iota
)

func (source SourceDialect) String() string {
	switch source {
	case SQLLITE_DIALECT:
		return appstrings.SQLLITE_DIALECT
	default:
		return appstrings.UNKNOWN_DIALECT
	}
}

type StoreConfig struct {
	StoreType     StoreType
	SourceType    SourceType
	SourceDialect SourceDialect
}
