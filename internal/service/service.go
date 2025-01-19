package service

import (
	"context"

	"github.com/YnMann/go-sister-design-site/internal/config"
	"github.com/YnMann/go-sister-design-site/internal/repository"
)

type Service struct {
}

func New(ctx context.Context, cfg *config.Config, repo *repository.Repository) *Service {
	return &Service{}
}
