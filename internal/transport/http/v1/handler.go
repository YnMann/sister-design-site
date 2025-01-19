package v1

/*
Package v1
Пакет предоставляет первую версию api для
*/

import (
	"github.com/YnMann/go-sister-design-site/internal/config"
	"github.com/YnMann/go-sister-design-site/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	cfg     *config.Config
}

func New(service *service.Service, cfg *config.Config) *Handler {
	return &Handler{
		service: service,
		cfg:     cfg,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	// v1 := api.Group("/v1")
	// {

	// }
}
