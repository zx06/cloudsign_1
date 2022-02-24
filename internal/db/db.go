package db

import (
	"context"

	"github.com/zx06/cloudsign/internal/provider"
)

type DB interface {
	CreateConfig(ctx context.Context, cfg *provider.SignConfig) error
	GetConfig(ctx context.Context, name string) (*provider.SignConfig, error)
	UpdateConfig(ctx context.Context, cfg *provider.SignConfig) error
	DeleteConfig(ctx context.Context, name string) error
	ListConfig(ctx context.Context) ([]provider.SignConfig, error)
}