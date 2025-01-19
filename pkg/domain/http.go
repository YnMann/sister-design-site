package domain

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/YnMann/go-sister-design-site/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ResponseWithMessage struct {
	Message string `json:"message"`
}

type ResponseMessageWithID struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

// Api response for Ping request
type Pong struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func GetBearerToken(c *gin.Context) (string, error) {
	header := c.GetHeader(AuthorizationHeader)
	if header == EmptyString {
		NewErrorResponse(c, http.StatusUnauthorized, ErrEmptyAuthHeader.Error(), ErrEmptyAuthHeader)
		return EmptyString, ErrEmptyAuthHeader
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, ErrInvalidAuthHeader.Error(), ErrInvalidAuthHeader)
		return EmptyString, ErrInvalidAuthHeader
	}

	return headerParts[1], nil
}

func GetAuthToken(c *gin.Context) (string, error) {
	token := c.GetHeader(AuthorizationHeader)
	if token == EmptyString {
		NewErrorResponse(c, http.StatusUnauthorized, ErrEmptyAuthHeader.Error(), ErrEmptyAuthHeader)
		return EmptyString, ErrEmptyAuthHeader
	}

	return token, nil
}

func GetIpFromHeader(c *gin.Context) (string, error) {
	//get ip from headers
	ip := c.GetHeader("x-original-forwarded-for")

	if ip == EmptyString {
		ip = c.GetHeader("x-forwarded-for")
	}

	if ip == EmptyString {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidHeader.Error(), errors.New("empty header x-original-forwarded-for and x-forwarded-for"))
		return EmptyString, ErrEmptyIp
	}

	return ip, nil
}

// GetQueryParam returns query parameter value.
func GetQueryParam[T any](c *gin.Context, name string) (T, error) {
	queryValue := c.Query(name)

	if queryValue == "" {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), ErrMissingQueryParam)
		var zeroValue T
		return zeroValue, ErrMissingQueryParam
	}

	var result T

	switch any(result).(type) {
	case int:
		parsedValue, err := strconv.Atoi(queryValue)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
			return result, err
		}
		result = any(parsedValue).(T)
	case int64:
		parsedValue, err := strconv.ParseInt(queryValue, 10, 64)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
			return result, err
		}
		result = any(parsedValue).(T)
	case float64:
		parsedValue, err := strconv.ParseFloat(queryValue, 64)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
			return result, err
		}
		result = any(parsedValue).(T)
	case bool:
		parsedValue, err := strconv.ParseBool(queryValue)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
			return result, err
		}
		result = any(parsedValue).(T)
	case string:
		result = any(queryValue).(T)
	default:
		err := errors.New("unsupported type")
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
		return result, err
	}

	return result, nil
}

// GetPathParam returns path parameter value.
func GetPathParam[T any](c *gin.Context, name string) (T, error) {
	paramValue := c.Param(name)

	if paramValue == "" {
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), ErrMissingQueryParam)
		var zeroValue T
		return zeroValue, ErrMissingQueryParam
	}

	var result T

	switch any(result).(type) {
	case int:
		parsedValue, err := strconv.Atoi(paramValue)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
			return result, err
		}
		result = any(parsedValue).(T)
	case int64:
		parsedValue, err := strconv.ParseInt(paramValue, 10, 64)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
			return result, err
		}
		result = any(parsedValue).(T)
	case float64:
		parsedValue, err := strconv.ParseFloat(paramValue, 64)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
			return result, err
		}
		result = any(parsedValue).(T)
	case bool:
		parsedValue, err := strconv.ParseBool(paramValue)
		if err != nil {
			NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
			return result, err
		}
		result = any(parsedValue).(T)
	case string:
		result = any(paramValue).(T)
	default:
		err := errors.New("unsupported type")
		NewErrorResponse(c, http.StatusBadRequest, ErrInvalidQueryParameter.Error(), err)
		return result, err
	}

	return result, nil
}

// ResponseWithMessage logs message and responds to client with provided status code and messageForUser
func ResponseWithFile(c *gin.Context, file []byte, filename string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Length", strconv.Itoa(len(file)))

	c.Data(http.StatusOK, "application/octet-stream", file)
}

// NewCouldntProcessResp loggs message and responds to client with provided status code and message "couldn`t process request"
func NewCouldntProcessResp(c *gin.Context, statusCode int, message string) {
	logger.Error(fmt.Sprintf("error: %s", message))
	c.AbortWithStatusJSON(statusCode, ResponseWithMessage{CouldntProcessRequest.Error()})
}
