package repositories

type IRepository interface {
	Any()
}

type IRepositoryProvider interface {
	GetUserRepository() UserRepository
}
