package config

type AppConfig struct {
	RouterManagerConfig  RouterManagerConfig
	ServiceManagerConfig ServiceManagerConfig
	StoreManagerConfig StoreManagerConfig
}

func DefaultAppConfig() *AppConfig {
	return &AppConfig{}
}
