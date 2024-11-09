package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/ELRAS1/auth/internal/model"
	"github.com/ELRAS1/auth/internal/repository/auth/converter"
	modulRepo "github.com/ELRAS1/auth/internal/repository/auth/model"
	"github.com/ELRAS1/auth/internal/repository/auth/utils"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	const nm = "[RepoCreate]"

	hashedPassword, err := utils.EncryptedPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	query, args, err := sq.Insert(dbName).
		Columns(name, email, password, role, createdAt).
		Values(req.Name, req.Email, hashedPassword, req.Role, time.Now()).
		Suffix(returningID).
		PlaceholderFormat(sq.Dollar).
		ToSql()
		
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	resp := &modulRepo.CreateResponse{}
	c, cancel := context.WithTimeout(ctx, time.Second*deadline)
	defer cancel()

	conn, err := r.Db.Acquire(c)
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}
	defer conn.Release()

	err = conn.QueryRow(c, query, args...).Scan(&resp.Id)
	if err != nil {
		return nil, fmt.Errorf("%s %w", nm, err)
	}

	return converter.RepoCreateToModel(resp), nil
}
