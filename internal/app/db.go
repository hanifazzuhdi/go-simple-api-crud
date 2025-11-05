package app

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetDbConnection(ctx context.Context, dbUrl string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, dbUrl)

	if err != nil {
		panic(err)
	}

	return conn
}
