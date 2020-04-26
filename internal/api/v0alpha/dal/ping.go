package dal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPong godoc
// @Summary Get pong
// @Description display pong
// @Produce  json
// @Success 200 {string} string "pong"
// @Router /v0alpha/ping [get]
func GetPong() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}
}
