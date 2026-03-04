package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"url-shortener/internal/config"
	"url-shortener/internal/storage"
)

func TestShortenAndResolve(t *testing.T) {
	gin.SetMode(gin.TestMode)

	store := storage.NewMemoryStore()
	cfg := config.Config{RateLimitEnabled: false, RateLimitPerMin: 10, RateLimitWindowS: 60}
	server := NewServer(store, cfg)

	router := gin.New()
	server.RegisterRoutes(router)

	payload := map[string]string{"url": "https://example.com/path"}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", recorder.Code)
	}

	var response struct {
		ShortURL string `json:"short_url"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if response.ShortURL == "" {
		t.Fatalf("expected short_url")
	}

	shortPath := response.ShortURL
	if idx := bytes.LastIndex([]byte(shortPath), []byte("/")); idx >= 0 {
		shortPath = shortPath[idx:]
	}
	if shortPath == response.ShortURL {
		shortPath = "/" + shortPath
	}

	resolveReq := httptest.NewRequest(http.MethodGet, shortPath, nil)
	resolveRecorder := httptest.NewRecorder()
	router.ServeHTTP(resolveRecorder, resolveReq)

	if resolveRecorder.Code != http.StatusMovedPermanently {
		t.Fatalf("expected 301, got %d", resolveRecorder.Code)
	}
	if location := resolveRecorder.Header().Get("Location"); location != "https://example.com/path" {
		t.Fatalf("unexpected location: %s", location)
	}
}
