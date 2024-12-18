package config

type RouterType int

const (
	GIN RouterType = iota
	GRPC
)

func (routerType RouterType) String() string {
	switch routerType {
	case GIN:
		return "GIN"
	case GRPC:
		return "GRPC"
	default:
		return "UnKNOWN"
	}
}

type RouterConfig struct {
	Type RouterType
	Port string
}

func MockRouterGinConfig() *RouterConfig {
	return &RouterConfig{}
}
