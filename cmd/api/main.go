package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/alexbednarczyk/crud-gopher/docs"
	v0AlphaRoutes "github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/routes"
	common "github.com/alexbednarczyk/crud-gopher/internal/common"
	"github.com/alexbednarczyk/crud-gopher/internal/database/config"
	"github.com/alexbednarczyk/crud-gopher/internal/database/migrations"
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

	d, err := iofs.New(migrations.FS, ".")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		d,
		config.BuildMigrationURLConnectionString(e))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Warn(err)
	}

	r := gin.Default()

	// CORS configuration for local development
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.NoRoute(func(c *gin.Context) {})
	v0AlphaRoutes.SetUpRouter(r, connPool)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Debug("Start API Service")

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
