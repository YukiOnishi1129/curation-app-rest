package database

import (
	"context"
	"fmt"
	"os"

	"github.com/YukiOnishi1129/curation-app-rest/backend/internal/infrastructure/persistence"
	"github.com/jackc/pgx/v5"
)

func Init(ctx context.Context) (queries *persistence.Queries, err error) {
	conn, err := connectDB(ctx)
	if err != nil {
		println("Failed to connect to db init")
		return nil, err
	}
	queries = persistence.New(conn)

	return queries, nil
}

func connectDB(ctx context.Context) (*pgx.Conn, error) {
	dsn := fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"))
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		println("Failed to connect to db")
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}
	return conn, nil
}
