package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"

	v0AlphaDAL "github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/dal"
)

// SetUpRouter sets up the routing logic for the defined API paths
func SetUpRouter(r *gin.Engine, db *pgxpool.Pool) {

	v0Alpha := r.Group("v0alpha")
	v0Alpha.GET("/ping", v0AlphaDAL.GetPong())

	v0AlphaExamples := v0Alpha.Group("/examples")
	v0AlphaExamples.GET("/attribute", v0AlphaDAL.Attribute())
	v0AlphaExamples.GET("/calc", v0AlphaDAL.CalcSumOfTwo())
	v0AlphaExamples.GET("/groups/:group_id/accounts/:account_id", v0AlphaDAL.GetAccountIDInGroup())
	v0AlphaExamples.GET("/header", v0AlphaDAL.UseHeader())
	v0AlphaExamples.GET("/securities", v0AlphaDAL.Securities())

	v0AlphaCRUD := v0Alpha.Group("/crud")
	v0AlphaCRUD.POST("/users", v0AlphaDAL.CreateUser(db))
	v0AlphaCRUD.GET("/users/:userId", v0AlphaDAL.GetSpecificUser(db))
	v0AlphaCRUD.GET("/users", v0AlphaDAL.GetAllUsers(db))
	v0AlphaCRUD.PATCH("/users/:userId", v0AlphaDAL.UpdateUser(db))
	v0AlphaCRUD.PUT("/users/:userId", v0AlphaDAL.ReplaceUser(db))
	v0AlphaCRUD.DELETE("/users/:userId", v0AlphaDAL.DeleteUser(db))

}
