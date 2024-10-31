package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/ELRAS1/auth/internal/model"
	"github.com/ELRAS1/auth/internal/repository/auth/converter"
	modulRepo "github.com/ELRAS1/auth/internal/repository/auth/model"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error) {
	const nm = "[RepoGet]"

	query, args, err := sq.Select(id, name, email, role, createdAt, updatedAt).From(dbName).
		Where(sq.Eq{id: req.Id}).
		PlaceholderFormat(sq.Dollar).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	c, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	conn, err := r.Db.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}
	defer conn.Release()

	rows, err := conn.Query(c, query, args...)

	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}
	defer rows.Close()

	resp := modulRepo.GetResponse{}
	if rows.Next() {
		err := rows.Scan(&resp.Id, &resp.Name, &resp.Email, &resp.Role, &resp.CreatedAt, &resp.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s %w", nm, err)
		}
	} else {
		return nil, fmt.Errorf("%s no rows found", nm)
	}

	return converter.RepoGetToModel(&resp), nil
}
