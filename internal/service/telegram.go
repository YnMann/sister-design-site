package service

import (
	"github.com/YnMann/go-sister-design-site/internal/config"
	"github.com/YnMann/go-sister-design-site/internal/repository"
)

type Telegram interface {
}

type TelegramService struct {
	cfg  *config.Config
	repo *repository.Repository
}

func NewTelegramService(cfg *config.Config, repo *repository.Repository) *TelegramService {
	return &TelegramService{cfg: cfg, repo: repo}
}
