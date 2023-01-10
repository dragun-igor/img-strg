package service

import "time"

//go:generate mockgen -destination=mocks/mock.go -package=mocks . Storage

type Storage interface {
	SetBirthTimeFile(string, time.Time) error
	GetBirthTimeFile(string) (time.Time, error)
}
