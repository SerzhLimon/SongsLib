package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/SerzhLimon/SongsLib/config"
	_ "github.com/SerzhLimon/SongsLib/docs"
	serv "github.com/SerzhLimon/SongsLib/internal/transport"
	"github.com/SerzhLimon/SongsLib/pkg/postgres"
	"github.com/SerzhLimon/SongsLib/pkg/postgres/migrations"
)

//	@title			Songs Library
//	@version		1.0
//	@description	This is a simple songs library server.
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
func main() {

	gin.SetMode(gin.ReleaseMode)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Info("Loading configuration...")
	config := config.LoadConfig()
	logrus.Debugf("Configuration loaded: %+v", config)

	logrus.Info("Initializing PostgreSQL client...")
	db, err := postgres.InitPostgresClient(config.Postgres)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize PostgreSQL client")
	}
	defer func() {
		logrus.Info("Closing PostgreSQL connection...")
		db.Close()
		logrus.Info("PostgreSQL connection closed")
	}()
	logrus.Info("PostgreSQL client initialized successfully")

	logrus.Info("Running migrations...")
	err = migrations.Up(db)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to apply migrations")
	}
	logrus.Info("Migrations applied successfully")

	logrus.Info("Initializing server...")
	server := serv.NewServer(db)
	routes := serv.ApiHandleFunctions{
		Server: *server,
	}

	logrus.Info("Setting up router...")
	router := serv.NewRouter(routes)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logrus.Infof("Starting server on port %s...", ":8080")
	if err := router.Run(":8080"); err != nil {
		logrus.WithError(err).Fatal("Failed to start server")
	}
	logrus.Info("Server started successfully")
}
