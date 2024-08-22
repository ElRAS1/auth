package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/ELRAS1/auth/app/utils"
	"github.com/ELRAS1/auth/pkg/userApi"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Storage) SaveUser(ctx context.Context, req *userApi.CreateRequest) (int64, error) {
	const nm = "[SaveUser]"
	hashedPassw, err := utils.EncryptedPassw(req.Password)
	if err != nil {
		return 0, fmt.Errorf("%s %v", nm, err)
	}
	query, args, err := sq.Insert("users").
		Columns("name", "email", "password", "role", "created_at", "updated_at").
		Values(req.Name, req.Email, hashedPassw, req.Role, time.Now(), time.Now()).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return 0, fmt.Errorf("%s %v", nm, err)
	}
	var id int64
	parentCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	err = s.Db.QueryRowContext(parentCtx, query, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s %v", nm, err)
	}
	return id, nil
}

func (s *Storage) DeleteUser(ctx context.Context, req *userApi.DeleteRequest) error {
	const nm = "[DeleteUser]"
	query, args, err := sq.Delete("users").
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	parentCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	_, err = s.Db.ExecContext(parentCtx, query, args...)
	if err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	return nil
}

func (s *Storage) GetUsers(ctx context.Context, req *userApi.GetRequest) (*userApi.GetResponse, error) {
	const nm = "[GetUsers]"
	query, args, err := sq.Select("id", "name", "email", "role", "created_at", "updated_at").From("users").
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s %v", nm, err)
	}
	parentCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	resp := userApi.GetResponse{}
	rows, err := s.Db.QueryContext(parentCtx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s %v", nm, err)
	}
	defer rows.Close()
	var createdAt, updatedAt time.Time
	if rows.Next() {
		err := rows.Scan(&resp.Id, &resp.Name, &resp.Email, &resp.Role, &createdAt, &updatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s %v", nm, err)
		}
		resp.CreatedAt = timestamppb.New(createdAt)
		resp.UpdatedAt = timestamppb.New(updatedAt)
	} else {
		return nil, fmt.Errorf("%s no rows found", nm)
	}
	return &resp, nil
}

func (s *Storage) UpdateUser(ctx context.Context, req *userApi.UpdateRequest) error {
	const nm = "[UpdateUser]"
	query, args, err := sq.Select("id").From("users").Where(sq.Eq{"id": req.Id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	var id int64
	parentCtx, cancel := context.WithTimeout(ctx, time.Second*15)
	defer cancel()
	err = s.Db.QueryRowContext(parentCtx, query, args...).Scan(&id)
	if err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	sql := sq.Update("users").Set("updated_at", time.Now()).Where(sq.Eq{"id": req.Id}).PlaceholderFormat(sq.Dollar)
	if req.Name.Value != "" {
		sql = sql.Set("name", req.Name.Value)
	}
	if req.Email.Value != "" {
		sql = sql.Set("email", req.Email.Value)
	}
	query, args, err = sql.ToSql()
	if err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	_, err = s.Db.ExecContext(parentCtx, query, args...)
	if err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	return nil
}
