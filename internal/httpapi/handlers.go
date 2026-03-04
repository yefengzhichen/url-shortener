package httpapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"url-shortener/internal/config"
	"url-shortener/internal/ratelimit"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"
)

type Server struct {
	store   storage.Store
	limiter *ratelimit.Limiter
}

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortURL string `json:"short_url"`
}

func NewServer(store storage.Store, cfg config.Config) *Server {
	var limiter *ratelimit.Limiter
	if cfg.RateLimitEnabled {
		limiter = ratelimit.New(cfg.RateLimitPerMin, time.Duration(cfg.RateLimitWindowS)*time.Second)
	}
	return &Server{store: store, limiter: limiter}
}

func (server *Server) RegisterRoutes(router *gin.Engine) {
	router.Use(ErrorMiddleware())
	router.POST("/shorten", server.handleShorten)
	router.GET("/:short", server.handleResolve)
	router.GET("/healthz", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}

func (server *Server) handleShorten(ctx *gin.Context) {
	if server.limiter != nil {
		clientIP := ctx.ClientIP()
		if !server.limiter.Allow(clientIP) {
			respondError(ctx, http.StatusTooManyRequests, "rate limit exceeded")
			return
		}
	}

	var req shortenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respondError(ctx, http.StatusBadRequest, "invalid json")
		return
	}
	if req.URL == "" {
		respondError(ctx, http.StatusBadRequest, "url is required")
		return
	}
	if err := shortener.ValidateURL(req.URL); err != nil {
		respondError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	code, err := server.createUniqueCode(ctx, req.URL)
	if err != nil {
		respondError(ctx, http.StatusInternalServerError, "failed to create short url")
		return
	}

	shortURL := shortener.BuildShortURL(ctx.Request, code)
	ctx.JSON(http.StatusOK, shortenResponse{ShortURL: shortURL})
}

func (server *Server) handleResolve(ctx *gin.Context) {
	code := ctx.Param("short")
	if code == "" {
		respondError(ctx, http.StatusBadRequest, "short code is required")
		return
	}

	target, err := server.store.Get(ctx.Request.Context(), code)
	if err != nil {
		if _, ok := err.(storage.NotFoundError); ok {
			respondError(ctx, http.StatusNotFound, "short code not found")
			return
		}
		respondError(ctx, http.StatusInternalServerError, "failed to resolve short code")
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, target)
}

func (server *Server) createUniqueCode(ctx context.Context, target string) (string, error) {
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}
		code, err := shortener.GenerateCode()
		if err != nil {
			return "", err
		}
		ok, err := server.store.SetIfAbsent(ctx, code, target)
		if err != nil {
			return "", err
		}
		if ok {
			return code, nil
		}
	}
}

func respondError(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{"error": message})
}
