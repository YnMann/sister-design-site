package repository

import "github.com/YnMann/go-sister-design-site/internal/config"

type Repository struct {
}

func New(cfg *config.Config) *Repository {
	return &Repository{}
}
