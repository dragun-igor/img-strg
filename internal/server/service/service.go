package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"
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

// Image storage service
type Service struct {
	strg.ImageStorageServer
	db               Storage
	storagePath      string
	limitLoadCounter int64
	limitListCounter int64
}

// Checking dir exists
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// Getting new image storage service
func New(db Storage, storagePath string) (*Service, error) {
	if !dirExists(storagePath) {
		if err := os.Mkdir(storagePath, 0755); err != nil {
			return nil, err
		}
	}
	return &Service{
		db:          db,
		storagePath: storagePath,
	}, nil
}

// Checking file exists
func fileExists(path string) bool {
	info, err := os.Stat(path)
	fmt.Println(info)
	fmt.Println(err)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Sending image to storage
func (s *Service) SendImage(ctx context.Context, r *strg.SendImageRequest) (*emptypb.Empty, error) {
	atomic.AddInt64(&s.limitLoadCounter, 1)
	defer func() {
		atomic.AddInt64(&s.limitLoadCounter, -1)
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
	_, err = file.Write(r.GetImage())
	if err != nil {
		return nil, convert(err)
	}
	return &emptypb.Empty{}, nil
}

// Getting image by filename
func (s *Service) GetImage(ctx context.Context, r *strg.GetImageRequest) (*strg.GetImageResponse, error) {
	atomic.AddInt64(&s.limitLoadCounter, 1)
	defer func() {
		atomic.AddInt64(&s.limitLoadCounter, -1)
	}()
	file, err := os.Open(s.storagePath + r.GetName())
	if err != nil {
		return nil, convert(err)
	}
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, convert(err)
	}
	return &strg.GetImageResponse{Image: b}, nil
}

// Getting images data in storage folder (filename, creation time, modification time)
func (s *Service) GetImagesList(ctx context.Context, r *emptypb.Empty) (*strg.GetImagesListResponse, error) {
	atomic.AddInt64(&s.limitListCounter, 1)
	defer func() {
		atomic.AddInt64(&s.limitListCounter, -1)
	}()
	files, err := os.ReadDir(s.storagePath)
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
