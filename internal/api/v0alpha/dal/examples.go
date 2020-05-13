package dal

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

// Attribute godoc
// @Summary attribute example
// @Description attribute
// @Tags examples
// @Accept json
// @Produce json
// @Param enumstring query string false "string enums" Enums(A, B, C)
// @Param enumint query int false "int enums" Enums(1, 2, 3)
// @Param enumnumber query number false "int enums" Enums(1.1, 1.2, 1.3)
// @Param string query string false "string valid" minlength(5) maxlength(10)
// @Param int query int false "int valid" mininum(1) maxinum(10)
// @Param default query string false "string default" default(A)
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /v0alpha/examples/attribute [get]
func Attribute() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("enumstring=%s enumint=%s enumnumber=%s string=%s int=%s default=%s",
			c.Query("enumstring"),
			c.Query("enumint"),
			c.Query("enumnumber"),
			c.Query("string"),
			c.Query("int"),
			c.Query("default"),
		))
	}
}

// CalcSumOfTwo godoc
// @Summary calc example
// @Description plus
// @Tags examples
// @Accept json
// @Produce json
// @Param val1 query int true "used for calc"
// @Param val2 query int true "used for calc"
// @Success 200 {integer} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /v0alpha/examples/calc [get]
func CalcSumOfTwo() func(c *gin.Context) {
	return func(c *gin.Context) {
		val1, err := strconv.Atoi(c.Query("val1"))
		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		val2, err := strconv.Atoi(c.Query("val2"))
		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		ans := val1 + val2
		c.String(http.StatusOK, "%d", ans)
	}
}

// GetAccountIDInGroup godoc
// @Summary path params example
// @Description path params
// @Tags examples
// @Accept json
// @Produce json
// @Param group_id path int true "Group ID"
// @Param account_id path int true "Account ID"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /v0alpha/examples/groups/{group_id}/accounts/{account_id} [get]
func GetAccountIDInGroup() func(c *gin.Context) {
	return func(c *gin.Context) {
		groupID, err := strconv.Atoi(c.Param("group_id"))
		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		accountID, err := strconv.Atoi(c.Param("account_id"))
		if err != nil {
			httputil.NewError(c, http.StatusBadRequest, err)
			return
		}
		c.String(http.StatusOK, "group_id=%d account_id=%d", groupID, accountID)
	}
}

// UseHeader godoc
// @Summary custome header example
// @Description custome header
// @Tags examples
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /v0alpha/examples/header [get]
func UseHeader() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, c.GetHeader("Authorization"))
	}
}

// Securities godoc
// @Summary custome header example
// @Description custome header
// @Tags examples
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {string} string "answer"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Security ApiKeyAuth
// @Security OAuth2Implicit[admin, write]
// @Router /v0alpha/examples/securities [get]
func Securities() func(c *gin.Context) {
	return func(c *gin.Context) {

	}
}
