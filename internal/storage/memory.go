package storage

import (
	"context"
	"sync"
)

type MemoryStore struct {
	data sync.Map
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

func (store *MemoryStore) Get(_ context.Context, code string) (string, error) {
	value, ok := store.data.Load(code)
	if !ok {
		return "", NotFoundError{Code: code}
	}
	target, ok := value.(string)
	if !ok {
		return "", NotFoundError{Code: code}
	}
	return target, nil
}

func (store *MemoryStore) Set(_ context.Context, code string, target string) error {
	store.data.Store(code, target)
	return nil
}

func (store *MemoryStore) SetIfAbsent(_ context.Context, code string, target string) (bool, error) {
	_, loaded := store.data.LoadOrStore(code, target)
	return !loaded, nil
}
