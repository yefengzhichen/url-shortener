package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"url-shortener/internal/config"
	"url-shortener/internal/httpapi"
	"url-shortener/internal/storage"
)

func main() {
	cfg := config.Load()

	store, cleanup, err := buildStore(cfg)
	if err != nil {
		log.Fatalf("failed to initialize storage: %v", err)
	}
	if cleanup != nil {
		defer cleanup()
	}

	router := gin.New()
	router.Use(gin.Recovery())

	server := httpapi.NewServer(store, cfg)
	server.RegisterRoutes(router)

	addr := ":8080"
	if value := os.Getenv("PORT"); value != "" {
		addr = ":" + value
	}

	if err := router.Run(addr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func buildStore(cfg config.Config) (storage.Store, func(), error) {
	if cfg.RedisURL == "" {
		return storage.NewMemoryStore(), nil, nil
	}

	options, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		return nil, nil, err
	}
	client := redis.NewClient(options)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		_ = client.Close()
	}
	return storage.NewRedisStore(client), cleanup, nil
}
