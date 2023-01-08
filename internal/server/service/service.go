package service

import (
	"context"
	"io"
	"log"
	"os"
	"sync"
	"syscall"
	"time"

	strg "github.com/dragun-igor/img-strg/proto/api"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	limitLoad int64 = 10
	limitList int64 = 100
)

type Service struct {
	strg.ImageStorageServer
	mu               *sync.Mutex
	db               Storage
	storagePath      string
	limitLoadCounter int64
	limitLoadCond    *sync.Cond
	limitListCounter int64
	limitListCond    *sync.Cond
}

// Проверка существования папки
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// Конструктор сервиса
func New(db Storage, storagePath string) (*Service, error) {
	if !dirExists(storagePath) {
		if err := os.Mkdir(storagePath, 0755); err != nil {
			return nil, err
		}
	}
	return &Service{
		db:            db,
		storagePath:   storagePath,
		mu:            &sync.Mutex{},
		limitLoadCond: sync.NewCond(&sync.Mutex{}),
		limitListCond: sync.NewCond(&sync.Mutex{}),
	}, nil
}

// Проверка существования файла
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Отправка изображения в бинарном виде
func (s *Service) SendImage(ctx context.Context, r *strg.SendImageRequest) (*emptypb.Empty, error) {
	s.limitLoadCond.L.Lock()
	for s.limitLoadCounter >= limitLoad {
		s.limitLoadCond.Wait()
	}
	s.limitLoadCounter++
	s.limitLoadCond.L.Unlock()

	defer func() {
		s.limitLoadCond.L.Lock()
		s.limitLoadCounter--
		s.limitLoadCond.L.Unlock()
		s.limitLoadCond.Signal()
	}()

	var file *os.File
	var err error

	if !fileExists(s.storagePath + r.GetName()) {
		file, err = os.Create(s.storagePath + r.GetName())
		if err != nil {
			return nil, convert(err)
		}
		if err := s.db.SetBirthTimeFile(r.GetName(), time.Now()); err != nil {
			return nil, convert(err)
		}
	} else {
		file, err = os.OpenFile(s.storagePath+r.GetName(), os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return nil, convert(err)
		}
	}

	defer file.Close()
	s.mu.Lock()
	_, err = file.Write(r.GetImage())
	s.mu.Unlock()
	if err != nil {
		return nil, convert(err)
	}
	return &emptypb.Empty{}, nil
}

// Получение изображения в бинарном виде по имени файла
func (s *Service) GetImage(ctx context.Context, r *strg.GetImageRequest) (*strg.GetImageResponse, error) {
	s.limitLoadCond.L.Lock()
	for s.limitLoadCounter >= limitLoad {
		s.limitLoadCond.Wait()
	}
	s.limitLoadCounter++
	s.limitLoadCond.L.Unlock()

	defer func() {
		s.limitLoadCond.L.Lock()
		s.limitLoadCounter--
		s.limitLoadCond.L.Unlock()
		s.limitLoadCond.Signal()
	}()
	file, err := os.Open(s.storagePath + r.GetName())
	if err != nil {
		return nil, convert(err)
	}
	s.mu.Lock()
	b, err := io.ReadAll(file)
	s.mu.Unlock()
	if err != nil {
		return nil, convert(err)
	}
	return &strg.GetImageResponse{Image: b}, nil
}

// Получение информации о файлах (название, вреям создания и время последней модификации)
func (s *Service) GetImagesList(ctx context.Context, r *emptypb.Empty) (*strg.GetImagesListResponse, error) {
	s.limitListCond.L.Lock()
	for s.limitListCounter >= limitList {
		s.limitListCond.Wait()
	}
	s.limitListCounter++
	s.limitListCond.L.Unlock()

	defer func() {
		s.limitListCond.L.Lock()
		s.limitListCounter--
		s.limitListCond.L.Unlock()
		s.limitListCond.Signal()
	}()
	s.mu.Lock()
	files, err := os.ReadDir(s.storagePath)
	s.mu.Unlock()
	if err != nil {
		return nil, convert(err)
	}
	images := make([]*strg.Images, 0, len(files))
	for _, dirEntry := range files {
		fileInfo, err := os.Stat(s.storagePath + dirEntry.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		stat := fileInfo.Sys().(*syscall.Stat_t)
		mtime := time.Unix(stat.Mtim.Sec, 0) //stat.Mtim.Nsec)
		ctime, err := s.db.GetBirthTimeFile(dirEntry.Name())
		if err != nil {
			log.Println(err)
			continue
		}
		images = append(images, &strg.Images{
			Name:             fileInfo.Name(),
			CreationTime:     timestamppb.New(ctime),
			ModificationTime: timestamppb.New(mtime),
		})
	}
	return &strg.GetImagesListResponse{Images: images}, nil
}
