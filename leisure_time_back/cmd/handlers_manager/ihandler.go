package handlersmanager

import repositoriesmanager "leisure_time_back/cmd/repositories_manager"

type IHandler interface {
	BindingRepository(
		repositoriesmanager.IRepository,
	)
}
