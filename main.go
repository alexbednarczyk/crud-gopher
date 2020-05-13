package main

import (
	"github.com/gin-gonic/gin"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/alexbednarczyk/crud-gopher/docs"
	v0AlphaRoutes "github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/routes"
	common "github.com/alexbednarczyk/crud-gopher/internal/common"
	"github.com/alexbednarczyk/crud-gopher/internal/database/config"
	"github.com/alexbednarczyk/crud-gopher/internal/environment/variables"
)

// @title Golang API Service
// @version 1.0
// @description This is an example Golang API service.
// @termsOfService http://swagger.io/terms/

// @license.name MIT
// @license.url https://github.com/alexbednarczyk/crud-gopher/blob/master/LICENSE

// @BasePath /
func main() {
	e := variables.FetchAllEnvironmentVariables()
	common.SetLogLevel(e.LogLevel)
	common.StandardFormatLogs(e.Deployed)

	connPool, err := config.GetPostgresDBConnectionPool(e)
	if err != nil {
		log.Fatal(err)
	}
	defer connPool.Close()

	m, err := migrate.New(
		"file://internal/database/migrations",
		config.BuildMigrationURLConnectionString(e))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Warn(err)
	}

	r := gin.Default()
	r.NoRoute(func(c *gin.Context) {})
	v0AlphaRoutes.SetUpRouter(r, connPool)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Debug("Start API Service")

	r.Run() //nolint
}
