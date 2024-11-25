package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbName    = "users"
	id        = "id"
	name      = "name"
	email     = "email"
	password  = "password"
	role      = "role"
	createdAt = "created_at"
	updatedAt = "updated_at"

	returningID = "RETURNING id"
	deadline    = 10 // context deadline
)

type repo struct {
	Db *pgxpool.Pool
}

func New(dbClient *pgxpool.Pool) *repo {
	return &repo{
		Db: dbClient,
	}
}
