package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"log/slog"
	"segments-api/internal/config"
	"segments-api/internal/controller/rest"
	"segments-api/internal/logger"
	"segments-api/internal/logger/sl"
	"segments-api/internal/repository"
	"segments-api/internal/service"
	"segments-api/pkg/database"
)

func Run(configDir string) {
	//configuration
	cfg := config.MustLoad(configDir)

	//init logger
	log := logger.SetupLogger(config.Env())

	//postgresql connection
	//conn := database.NewPostgresConnection(cfg.Postgres)
	conn := database.New(context.Background(), cfg.Postgres, log)

	//clean architecture: handler -> service -> repository

	//init repository
	segmentRepository := repository.New(conn, log)

	//init service
	segmentService := service.New(segmentRepository, log)

	//init handler
	route := rest.New(segmentService, log)

	//run server
	startServer(route, cfg, log)

	// Graceful Shutdown

}

func startServer(route rest.SegmentRoute, cfg *config.Config, log *slog.Logger) {
	srv := echo.New()
	srv.GET("/swagger/*any", echoSwagger.WrapHandler)

	srv.POST("/segments", route.Create)
	srv.DELETE("/segments", route.Delete)
	srv.POST("/segments/users", route.AddUser)
	//srv.POST("/segments/users/:userId", route.AddUser)
	srv.GET("/segments/users/:userId", route.GetAllByUser)

	err := srv.Start(fmt.Sprintf(":%d", cfg.Http.Port))
	if err != nil {
		log.Error("server starting failed", sl.Err(err))
	}
}
