package v1

import (
	"github.com/gin-gonic/gin"
)

// Inits v1 router group resource /tasks.
func (h *Handler) InitHealthRoute(api *gin.RouterGroup) {
	health := api.Group("/health")
	{
		health.GET("", h.healthStatus)
	}
}

// ShowAccount godoc
// @Summary      Api health check
// @Description  Returns server status
// @Tags         Health
// @Produce      json
// @Success      200
// @Router       /health [get].
func (h *Handler) healthStatus(c *gin.Context) {
	s := h.health.ReadStatus()
	c.JSON(200, gin.H{
		"Uptime":                              s["utime"],
		"Total heap allocations since deploy": s["memtot"],
		"GC metadata":                         s["gc"],
		"Total OS memory allocated":           s["sys"],
	})
}
