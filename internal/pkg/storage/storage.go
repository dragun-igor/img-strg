package storage

import (
	"strconv"
	"time"
)

type Storage struct {
	redis Redis
}

func New(redis Redis) *Storage {
	return &Storage{
		redis: redis,
	}
}

func (s *Storage) SetBirthTimeFile(key string, value time.Time) error {
	return s.redis.Set(key, value.Unix())
}

func (s *Storage) GetBirthTimeFile(key string) (time.Time, error) {
	data, err := s.redis.Get(key)
	if err != nil {
		return time.Time{}, err
	}
	tint, err := strconv.Atoi(string(data.([]uint8)))
	if err != nil {
		return time.Time{}, err
	}
	t := time.Unix(int64(tint), 0)
	return t, nil
}
