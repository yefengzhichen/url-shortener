package storage

import (
	"context"
	"testing"
)

func TestMemoryStoreGetSet(t *testing.T) {
	store := NewMemoryStore()
	ctx := context.Background()

	if err := store.Set(ctx, "abc", "https://example.com"); err != nil {
		t.Fatalf("set failed: %v", err)
	}

	value, err := store.Get(ctx, "abc")
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if value != "https://example.com" {
		t.Fatalf("unexpected value: %s", value)
	}
}

func TestMemoryStoreSetIfAbsent(t *testing.T) {
	store := NewMemoryStore()
	ctx := context.Background()

	ok, err := store.SetIfAbsent(ctx, "abc", "https://example.com")
	if err != nil || !ok {
		t.Fatalf("expected first set to succeed")
	}

	ok, err = store.SetIfAbsent(ctx, "abc", "https://example.com/2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ok {
		t.Fatalf("expected second set to be rejected")
	}
}
