package store

type Store interface {
	User() UserRepository
}

type RedisStore interface {
	Get(string) (string, error)
	Set(string, string) error
	Delete(string) error
}
