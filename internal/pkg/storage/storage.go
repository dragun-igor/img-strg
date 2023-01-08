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

// Сохраняет время создания файла в редис (в линуксе не хранится время созданяи файла)
func (s *Storage) SetBirthTimeFile(key string, value time.Time) error {
	return s.redis.Set(key, value.Unix())
}

// Получение времени создания файла
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
