package config

import appstrings "leisure_time/internal/domain/strings"

type RepositoryType int

const (
	UserRepositoryType RepositoryType = iota
)

func (repositoryType RepositoryType) String() string {
	switch repositoryType {
	case UserRepositoryType:
		return appstrings.USER_TYPE
	default:
		return appstrings.UNKNOWN
	}
}

type RepositoryConfig struct {
	Type RepositoryType
}
