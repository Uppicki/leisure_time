package config

type AppConfig struct {
	RouterManagerConfig  RouterManagerConfig
	ServiceManagerConfig ServiceManagerConfig
}

func DefaultAppConfig() *AppConfig {
	return &AppConfig{}
}
