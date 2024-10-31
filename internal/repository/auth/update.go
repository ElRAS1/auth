package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/ELRAS1/auth/internal/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Update(ctx context.Context, req *model.UpdateRequest) error {
	const nm = "[RepoUpdate]"

	c, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	sql := sq.Update(dbName).
		Set(updatedAt, time.Now()).
		Where(sq.Eq{id: req.Id}).
		PlaceholderFormat(sq.Dollar)

	if req.Name != "" {
		sql = sql.Set(name, req.Name)
	}

	if req.Email != "" {
		sql = sql.Set(email, req.Email)
	}

	query, args, err := sql.ToSql()
	if err != nil {
		return fmt.Errorf("%s %w", nm, err)
	}

	conn, err := r.Db.Acquire(c)
	if err != nil {
		return fmt.Errorf("%s %w", nm, err)
	}
	defer conn.Release()

	res, err := conn.Exec(c, query, args...)
	if err != nil {
		return fmt.Errorf("%s %w", nm, err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%s user was not found", nm)
	}

	return nil
}
