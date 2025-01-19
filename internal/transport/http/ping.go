package v1

import (
	"net/http"
	"time"

	"github.com/YnMann/go-sister-design-site/pkg/domain"
	"github.com/YnMann/go-sister-design-site/tools"
	"github.com/gin-gonic/gin"
)

// Inits v1 router group resource /tasks.
func (h *Handler) InitPingRoute(api *gin.RouterGroup) {
	ping := api.Group("/ping")
	{
		ping.GET("", h.IntiPingRoute)
	}
}

// ShowAccount godoc
// @Summary      Pong
// @Description  Returns pong object with server time (UTC)
// @Tags         Ping
// @Produce      json
// @Success      200 {object} domain.Pong
// @Router       /ping [get].
func (h *Handler) IntiPingRoute(c *gin.Context) {
	time := tools.ConvertTimeToJSTime(time.Now().UTC())

	c.JSON(http.StatusOK, domain.Pong{Message: "pong", Time: time})
}
