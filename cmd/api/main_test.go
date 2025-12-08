package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	v0AlphaRoutes "github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/routes"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := gin.Default()
	v0AlphaRoutes.SetUpRouter(router, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v0alpha/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
