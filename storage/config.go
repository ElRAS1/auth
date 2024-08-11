package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
}

func ConfigureStorage() (*Storage, error) {
	const nm = "[ConfigureStorage]"
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
	if err = db.Ping(); err != nil {
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
