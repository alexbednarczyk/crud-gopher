package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	v0AlphaRoutes "github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/routes"
	common "github.com/alexbednarczyk/crud-gopher/internal/common"
)

func TestPingRoute(t *testing.T) {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		common.NewError(c, http.StatusNotFound, errors.New("route not found"))
	})
	v0AlphaRoutes.SetUpRouter(router, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v0alpha/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestNoRoute(t *testing.T) {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		common.NewError(c, http.StatusNotFound, errors.New("route not found"))
	})
	v0AlphaRoutes.SetUpRouter(router, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/unknown", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "route not found")
}
