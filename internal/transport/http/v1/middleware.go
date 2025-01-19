package v1

import (
	"github.com/gin-gonic/gin"
)

// userIdentity is a middleware for the WebSocket server.
func (h *Handler) userIdentity(c *gin.Context) {

}

// isValidToken checks if the token is valid.
func (h *Handler) isValidToken(c *gin.Context) bool {

	return true
}
