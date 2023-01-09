package server

import (
	"context"
	"net"
	"os"

	"github.com/dragun-igor/img-strg/config"
	"github.com/dragun-igor/img-strg/internal/pkg/storage"
	"github.com/dragun-igor/img-strg/internal/server/resources"
	"github.com/dragun-igor/img-strg/internal/server/service"
	"google.golang.org/grpc"

	strg "github.com/dragun-igor/img-strg/proto/api"
)

type Server struct {
	grpc   *grpc.Server
	config *config.Config
	db     db
}

// Конструктор сервера
func New(cfg *config.Config) (*Server, error) {
	db, err := resources.InitRedis(cfg)
	if err != nil {
		return nil, err
	}
	storage := storage.New(db)
	grpc := grpc.NewServer([]grpc.ServerOption{}...)
	service, err := service.New(storage, cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	strg.RegisterImageStorageServer(grpc, service)
	return &Server{
		grpc:   grpc,
		config: cfg,
		db:     db,
	}, nil
}

func (s *Server) Serve(ctx context.Context) error {
	defer s.db.Close()
	lis, err := net.Listen("tcp", s.config.GRPCAddr)
	if err != nil {
		return err
	}
	sigCh := make(chan os.Signal, 1)
	go func() {
		<-sigCh
		s.grpc.GracefulStop()
	}()
	return s.grpc.Serve(lis)
}
