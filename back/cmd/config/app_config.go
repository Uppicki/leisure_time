package config

type AppConfig struct {
	RouterManagerConfig     RouterManagerConfig
	ServiceManagerConfig    ServiceManagerConfig
	RepositoryManagerConfig RepositoryManagerConfig
	StoreManagerConfig      StoreManagerConfig
}

func DefaultAppConfig() *AppConfig {
	return &AppConfig{}
}
