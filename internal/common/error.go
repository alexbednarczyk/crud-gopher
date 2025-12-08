package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewError example
func NewError(c *gin.Context, status int, err error) {
	if err == nil {
		err = errors.New(http.StatusText(status))
	}
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"internal server error"`
}
