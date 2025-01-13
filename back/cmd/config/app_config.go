package config

type AppMod int

const (
	Debug AppMod = iota
	Prod
)

type AppConfig struct {
	Mod                     AppMod
	RouterManagerConfig     RouterManagerConfig
	ServiceManagerConfig    ServiceManagerConfig
	RepositoryManagerConfig RepositoryManagerConfig
	StoreManagerConfig      StoreManagerConfig
}

func DefaultAppConfig() *AppConfig {
	return &AppConfig{}
}
