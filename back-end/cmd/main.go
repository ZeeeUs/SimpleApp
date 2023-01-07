package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZeeeUs/SimpleApp/internal/config"
	"github.com/ZeeeUs/SimpleApp/internal/service"
	bookStorage "github.com/ZeeeUs/SimpleApp/internal/storage"
	"github.com/ZeeeUs/SimpleApp/internal/transport/http"
	"github.com/ZeeeUs/SimpleApp/internal/transport/http/handlers"
	"github.com/jackc/pgx"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.NewConfig()
	log.Logger = config.Logger()

	connStr := connectionString(
		cfg.DbConfig.Host,
		cfg.DbConfig.Database,
		cfg.DbConfig.User,
		cfg.DbConfig.Password,
	)
	conn := newDbConnection(connStr)

	storage := bookStorage.New(conn)
	svc := service.New(log.Logger, storage)
	handler := handlers.New(log.Logger).WithBook(svc)
	server := http.NewServer(cfg.ServerConfig.Host).WithBookHandler(handler)

	corsSettings := http.CorsSettings(cfg.FrontConfig.Host)
	err := server.Run(corsSettings)
	if err != nil {
		log.Fatal().Err(err).Msg("server starting error")
	}
	log.Info().Msgf("http server starting at host %s", cfg.ServerConfig.Host)

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Info().Msg("http server shutdown")

	if err = server.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("server shutdown error")
	}
}

func connectionString(host, db, user, password string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		user,
		password,
		host,
		db,
	)
}

func newDbConnection(connStr string) *pgx.Conn {
	conf, err := pgx.ParseConnectionString(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse connection string")
	}
	conn, err := pgx.Connect(conf)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get db connection")
	}

	return conn
}
