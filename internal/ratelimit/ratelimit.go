package ratelimit

import (
	"sync"
	"time"
)

type Limiter struct {
	limit   int
	window  time.Duration
	mu      sync.Mutex
	clients map[string]*clientState
}

type clientState struct {
	count int
	reset time.Time
}

func New(limit int, window time.Duration) *Limiter {
	return &Limiter{
		limit:   limit,
		window:  window,
		clients: make(map[string]*clientState),
	}
}

func (limiter *Limiter) Allow(key string) bool {
	if limiter.limit <= 0 {
		return true
	}
	now := time.Now()

	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	state := limiter.clients[key]
	if state == nil || now.After(state.reset) {
		limiter.clients[key] = &clientState{
			count: 1,
			reset: now.Add(limiter.window),
		}
		return true
	}

	if state.count >= limiter.limit {
		return false
	}

	state.count++
	return true
}
