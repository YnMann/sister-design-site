package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/YnMann/go-sister-design-site/internal/config"
	"github.com/YnMann/go-sister-design-site/internal/repository"
	http_s "github.com/YnMann/go-sister-design-site/internal/server/http"
	"github.com/YnMann/go-sister-design-site/internal/service"
	http_h "github.com/YnMann/go-sister-design-site/internal/transport/http"
	"github.com/YnMann/go-sister-design-site/pkg/logger"
)

func main() {
	// init logger
	logger.ZapLoggerInit()

	// init context
	ctx, cancel := context.WithCancel(context.Background())

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	// cancel context on shutdown
	go func() {
		<-quit
		cancel()
	}()

	// Init config
	cfg, err := config.New(os.Getenv("IS_PROD"))
	if err != nil {
		logger.Fatalf("config: failed to load config, err: %v", err.Error())
	}

	// MARIADB
	// mdb, err := database.NewMariaDBConnection(cfg.MariaDB)
	// if err != nil {
	// 	logger.Fatalf("Failed to connect maria db, err - %v", err)
	// }
	// defer func(conn *database.MariaDBConnections) {
	// 	if err := conn.Default.Close(); err != nil {
	// 		logger.Errorf("Failed to close maria db, err - %v", err)
	// 	}
	// }(mdb)

	// init dep-s
	repo := repository.New(cfg)
	s := service.New(ctx, cfg, repo)
	hh := http_h.New(s, cfg)

	// init HTTP Server
	hsrv := http_s.New(cfg, hh.Init())

	// run http server
	go func(srv *http_s.Server) {
		if err = srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(fmt.Sprintf("error occurred while running http server: %s", err))
		}
	}(hsrv)

	logger.Info("go-sister-design-site is running")

	// wait for signal to stop
	<-ctx.Done()

	logger.Info("go-sister-design-site is stopped")
}
