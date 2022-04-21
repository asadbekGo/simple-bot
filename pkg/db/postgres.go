package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres drives

	"github.com/asadbekGo/telegram-message-service/config"
)

func ConnectionToPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlDbname,
		cfg.Postgres.PostgresqlPassword,
		cfg.Postgres.PostgresqlPort,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDb, nil
}
