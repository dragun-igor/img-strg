package main

import (
	"context"
	"log"

	"github.com/dragun-igor/img-strg/config"
	"github.com/dragun-igor/img-strg/internal/server"
)

func main() {
	cfg := config.Get()
	ctx := context.Background()
	s, err := server.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	if err := s.Serve(ctx); err != nil {
		log.Fatalln(err)
	}
}
