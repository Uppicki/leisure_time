package config

import appstrings "leisure_time/internal/domain/strings"

type ServiceType int

const (
	UserServiceType ServiceType = iota
)

func (serviceType ServiceType) String() string {
	switch serviceType {
	case UserServiceType:
		return appstrings.USER_TYPE
	default:
		return appstrings.UNKNOWN
	}
}

type ServiceConfig struct {
	Type ServiceType
}
