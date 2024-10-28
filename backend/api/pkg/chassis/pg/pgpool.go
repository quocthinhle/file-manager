package chassis_pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/quocthinhle/file-manager-api/pkg/chassis"
)

func NewPool(
	ctx context.Context,
) (conn *pgxpool.Pool, err error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		chassis.MustGetEnv("POSTGRES_USER"),
		chassis.MustGetEnv("POSTGRES_PASSWORD"),
		chassis.MustGetEnv("POSTGRES_HOST"),
		chassis.MustGetEnv("POSTGRES_PORT"),
		chassis.MustGetEnv("POSTGRES_NAME"),
	)

	var config *pgxpool.Config
	if config, err = pgxpool.ParseConfig(connectionString); err != nil {
		return
	}

	return pgxpool.NewWithConfig(ctx, config)
}
