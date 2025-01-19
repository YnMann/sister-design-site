package domain

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/YnMann/go-sister-design-site/pkg/logger"

	"github.com/gin-gonic/gin"
)

type JsonMessage struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

const ( // http strings
	UnauthorizedString string = "unauthorized"
	InvalidToken       string = "invalid token auth attempt"
	InvalidUser        string = "invalid user auth attempt"
	WrongContentType          = "wrong content type"
)

var ( // auth
	ErrEmptyAuthHeader   = errors.New("empty auth header")           // looks like - empty auth header
	ErrInvalidAuthHeader = errors.New("invalid auth header")         // looks like - invalid auth header
	ErrEmptyIp           = errors.New("ip from headers not founded") // looks like - ip from headers not founded
)

var ( // http
	ErrInvalidHeader         = errors.New("invalid header")            // looks like - invalid header
	ErrMissingQueryParam     = errors.New("missing query parameter")   // looks like - missing query parameter
	ErrInvalidQueryParameter = errors.New("invalid query parameter")   // looks like - invalid query parameter
	CouldntProcessRequest    = errors.New("could not process request") // looks like - could not process request
)

func AbortWithNotAllowedError(c *gin.Context, logMessage string) {
	logger.Error(fmt.Sprintf("error: %v", logMessage))
	c.AbortWithStatusJSON(http.StatusUnauthorized, JsonMessage{Message: UnauthorizedString})
}

// NewErrorResponse loggs message and responds to client with provided status code and messageForUser
func NewErrorResponse(c *gin.Context, statusCode int, messageForUser string, loggingErr error) {
	logger.Error(fmt.Sprintf("error: %v", loggingErr))
	c.AbortWithStatusJSON(statusCode, ResponseWithMessage{messageForUser})
}
