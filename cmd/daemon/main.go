package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/myrat012/test-work-song-lib/db"
	"github.com/myrat012/test-work-song-lib/internal/controller/http"
	"github.com/myrat012/test-work-song-lib/internal/migration"
	"github.com/myrat012/test-work-song-lib/internal/usecase"
	"github.com/myrat012/test-work-song-lib/pkg/config"
	"github.com/myrat012/test-work-song-lib/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// logger
	pwd, _ := os.Getwd()
	zLog, err := logger.New("info", filepath.Join(pwd, "logs"))
	if err != nil {
		zLog.Panic().Err(err).Msg("could not initialize Logger")
		return
	}
	zerolog.DefaultContextLogger = zLog
	log.Logger = *zLog

	// Load .env
	conf, err := config.LoadEnv(".env")
	if err != nil {
		return
	}

	// Create PostgreSQL connection string for pgx
	pool, err := db.NewPool(conf.DBConfig)
	if err != nil {
		log.Panic().Err(err).Msg("could not initialize db.NewPool")
		return
	}
	defer pool.Close()

	// Run migrations
	if err := migration.RunMigrations(conf.DBConfig.ConnString); err != nil {
		log.Panic().Err(err).Msg("could not initialize migration.RunMigrations")
		return
	}

	// Use cases
	useCases := usecase.LoadUseCases(pool)

	srv, err := http.NewService(*conf.APPConfig, useCases)
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.APPConfig.Host, conf.APPConfig.Port))
	if err != nil {
		panic(err)
	}
	log.Info().
		Str("listen", fmt.Sprintf("%s:%s", conf.APPConfig.Host, conf.APPConfig.Port)).
		Msg("Starting HTTP API Server")

	go func() {
		err = srv.Serve(listener)
		if err != nil {
			panic("HTTP server Error")
		}
	}()

	// Waiting signal
	signalChan := make(chan os.Signal, 1)
	quit := make(chan interface{})
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case <-quit:
			log.Warn().Msg("quit channel closed, closing listener")
			err = srv.Close()
			if err != nil {
				log.Error().Err(err).Msg("error during HTTP Server close")
			}
			err = listener.Close()
			if err != nil {
				log.Error().Err(err).Msg("error during TCP Listener close")
			}
			return
		case sig := <-signalChan:
			switch sig {
			case os.Interrupt, os.Kill, syscall.SIGTERM:
				log.Info().Msg("interrupt signal received, sending Quit signal")
				close(quit)
			default:
				log.Info().
					Any("signal", sig).
					Msg("signal received")
			}
		}
	}
}
