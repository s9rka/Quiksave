package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func InitDb() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_address := os.Getenv("DB_ADDRESS")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s", db_user, db_pass, db_address, db_name)

	var err error
	dbPool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		return fmt.Errorf("unable to create connection pool: %w", err)
	}
	fmt.Println("Database connection achieved!")
	return nil
}