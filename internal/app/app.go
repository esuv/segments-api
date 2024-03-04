package app

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
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
	//Configuration
	cfg := config.MustLoad(configDir)

	//Init logger
	log := logger.SetupLogger(config.Env())

	//Init repository layer
	//conn := database.NewPostgresConnection(cfg.Postgres)
	conn := database.New(context.TODO(), cfg.Postgres, log)
	segmentRepository := repository.New(conn, log)

	//Init services
	segmentService := service.New(segmentRepository, log)

	//Init handlers
	route := rest.New(segmentService)

	//Init server
	initServer(route, cfg, log)

	// Graceful Shutdown

}

func initServer(route rest.SegmentRoute, cfg *config.Config, log *slog.Logger) {
	srv := echo.New()

	srv.POST("/segments", route.Create)
	srv.DELETE("/segments", route.Delete)
	srv.POST("/segments/users/:userId", route.AddUser)
	srv.GET("/segments/users/:userId", route.GetAllByUser)

	err := srv.Start(fmt.Sprintf(":%d", cfg.Http.Port))
	if err != nil {
		log.Error("server starting failed", sl.Err(err))
	}
}
