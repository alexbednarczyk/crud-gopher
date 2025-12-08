package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"

	v0AlphaHandlers "github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/handlers"
)

// SetUpRouter sets up the routing logic for the defined API paths
func SetUpRouter(r *gin.Engine, db *pgxpool.Pool) {

	v0Alpha := r.Group("v0alpha")
	v0Alpha.GET("/ping", v0AlphaHandlers.GetPong())

	v0AlphaExamples := v0Alpha.Group("/examples")
	v0AlphaExamples.GET("/attribute", v0AlphaHandlers.Attribute())
	v0AlphaExamples.GET("/calc", v0AlphaHandlers.CalcSumOfTwo())
	v0AlphaExamples.GET("/groups/:group_id/accounts/:account_id", v0AlphaHandlers.GetAccountIDInGroup())
	v0AlphaExamples.GET("/header", v0AlphaHandlers.UseHeader())
	v0AlphaExamples.GET("/securities", v0AlphaHandlers.Securities())

	v0AlphaCRUD := v0Alpha.Group("/crud")
	v0AlphaCRUD.POST("/users", v0AlphaHandlers.CreateUser(db))
	v0AlphaCRUD.GET("/users/:userId", v0AlphaHandlers.GetSpecificUser(db))
	v0AlphaCRUD.GET("/users", v0AlphaHandlers.GetAllUsers(db))
	v0AlphaCRUD.PATCH("/users/:userId", v0AlphaHandlers.UpdateUser(db))
	v0AlphaCRUD.PUT("/users/:userId", v0AlphaHandlers.ReplaceUser(db))
	v0AlphaCRUD.DELETE("/users/:userId", v0AlphaHandlers.DeleteUser(db))

}
