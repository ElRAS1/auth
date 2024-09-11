package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
}

func ConfigureStorage(ctx context.Context) (*Storage, error) {
	const nm = "[ConfigureStorage]"

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err := load()
	if err != nil {
		return nil, fmt.Errorf("%s %v", nm, err)
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%s %v", nm, err)
	}

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("%s %v", nm, err)
	}

	return &Storage{Db: db}, nil
}

func load() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("no .env file found")
	}
	return nil
}
