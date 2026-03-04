package storage

import "context"

type Store interface {
	Get(ctx context.Context, code string) (string, error)
	Set(ctx context.Context, code string, target string) error
	SetIfAbsent(ctx context.Context, code string, target string) (bool, error)
}

type NotFoundError struct {
	Code string
}

func (err NotFoundError) Error() string {
	return "code not found"
}
