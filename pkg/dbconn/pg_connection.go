package dbconn

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type DB *pgxpool.Pool

func GetDBClient() (DB, error) {
	ctx := context.Background()
	dsn := os.Getenv("PSQL_URL")

	DB, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logrus.Fatalf("failed in database connection: %v", err.Error())
	}

	if err = DB.Ping(ctx); err != nil {
		panic(err)
	}

	return DB, nil
}
