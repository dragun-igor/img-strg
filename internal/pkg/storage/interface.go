package storage

//go:generate mockgen -destination=mocks/mock.go -package=mocks . Redis

type Redis interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error
}
