package api

import (
	"context"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/config"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/delivery/http"
	imageManager "github.com/denis-shcherbinin/gora-studio-test-task/internal/image-manager"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/repository"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/repository/sqlite"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/server"
	"github.com/denis-shcherbinin/gora-studio-test-task/internal/service"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title GORA Studio Photo Gallery Test Task API
// @version 1.0
// @description API Server for GORA Studio Photo Gallery Test Task

// @host localhost:8080
// @BasePath /api/v1/

// Run launches the API
func Run(configPath string) {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Config init
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		logrus.Fatalf("error occured while reading configs: %v", err)
	}

	// Repositories init
	db, err := sqlite.NewSqliteDb(cfg.SqliteConfig)
	if err != nil {
		logrus.Fatalf("error occured while initializing sqlite db: %v", err)
	}
	repo := repository.NewRepositories(db)

	// Services init
	services := service.NewServices(service.Deps{
		Repo:         repo,
		ImageManager: imageManager.NewPhotoUploader("./image", "/stat-img"),
	})

	// Handlers init
	handlers := http.NewHandler(services)

	// Server init
	srv := server.NewServer(cfg.HttpConfig, handlers.Init())
	go func() {
		if err = srv.Run(); err != nil {
			logrus.Fatalf("error occured while running http server: %v", err)
		}
	}()
	logrus.Print(" [___SERVER STARTED___] ")

	// Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err = srv.Shutdown(ctx); err != nil {
		logrus.Errorf("failed to shutdown server: %v", err)
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %v", err)
	}
}
