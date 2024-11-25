package user

import (
	"context"
	"fmt"
	"time"

	"github.com/ELRAS1/auth/internal/models/user/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Delete(ctx context.Context, req *model.DeleteRequest) error {
	const nm = "[RepoDelete]"

	query, args, err := sq.Delete(dbName).
		Where(sq.Eq{id: req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return fmt.Errorf("%s %w", nm, err)
	}

	c, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

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
