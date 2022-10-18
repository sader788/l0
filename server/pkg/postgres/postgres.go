package postgres

import (
	"WildberriesL0/server/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"time"
)

func ConnectDB(cfg *config.ConfigPostgres) (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.UserName, cfg.UserPassword, cfg.PostgresHost, cfg.PostgresPort, cfg.Database)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
