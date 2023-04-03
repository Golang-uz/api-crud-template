package postgres_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/storage"
	"github.com/realtemirov/api-crud-template/storage/postgres"
	"github.com/rs/zerolog"
)

var cfg *config.Config = &config.Config{
	PostgresHost:     "localhost",
	PostgresPort:     "4001",
	PostgresUser:     "postgres",
	PostgresPassword: "123456",
	PostgresDB:       "pack",
	PostgresSSLMode:  "disable",
}
var str storage.StorageI
var ctx context.Context = context.Background()
var logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
	Level(zerolog.InfoLevel).
	With().
	Timestamp().
	Caller().
	Logger()

func TestMain(m *testing.M) {

	psql, err := postgres.NewPostgres(ctx, cfg, logger)
	if err != nil {
		logger.Err(err)
	}

	str = psql

	os.Exit(m.Run())
}
