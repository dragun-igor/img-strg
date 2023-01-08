package service

import "time"

type Storage interface {
	SetBirthTimeFile(string, time.Time) error
	GetBirthTimeFile(string) (time.Time, error)
}
