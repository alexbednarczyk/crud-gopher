package routes

import (
	"github.com/gin-gonic/gin"

	v0AlphaDAL "github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal"
)

// SetUpRouter sets up the routing logic for the defined API paths
func SetUpRouter(r *gin.Engine) {

	v0Alpha := r.Group("v0alpha")
	v0Alpha.GET("/ping", v0AlphaDAL.GetPong())

	v0AlphaExamples := v0Alpha.Group("/examples")
	v0AlphaExamples.GET("/attribute", v0AlphaDAL.Attribute())
	v0AlphaExamples.GET("/calc", v0AlphaDAL.CalcSumOfTwo())
	v0AlphaExamples.GET("/groups/:group_id/accounts/:account_id", v0AlphaDAL.GetAccountIDInGroup())
	v0AlphaExamples.GET("/header", v0AlphaDAL.UseHeader())
	v0AlphaExamples.GET("/securities", v0AlphaDAL.Securities())

}
