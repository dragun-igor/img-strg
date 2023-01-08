package server

type db interface {
	Close() error
}
