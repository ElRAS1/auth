package storage

import (
	"fmt"
	"time"

	"github.com/ELRAS1/auth/pkg/userApi"
	sq "github.com/Masterminds/squirrel"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Storage) SaveUser(req *userApi.CreateRequest) (int64, error) {
	const nm = "[SaveUser]"
	query, sql, err := sq.Insert("users").
		Columns("name", "email", "password", "role", "created_at", "updated_at").
		Values(req.Name, req.Email, req.Password, req.Role, time.Now(), time.Now()).
		Suffix("RETURNING id").PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return 0, fmt.Errorf("%s %v", nm, err)
	}
	var id int64
	err = s.Db.QueryRow(query, sql...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s %v", nm, err)
	}
	return id, nil
}

func (s *Storage) DeleteUser(req *userApi.DeleteRequest) error {
	const nm = "[DeleteUser]"
	query, args, err := sq.Delete("users").
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	_, err = s.Db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("%s %v", nm, err)
	}
	return nil
}

func (s *Storage) GetUsers(req *userApi.GetRequest) (*userApi.GetResponse, error) {
	const nm = "[GetUsers]"
	query, args, err := sq.Select("id", "name", "email", "role", "created_at", "updated_at").From("users").
		Where(sq.Eq{"id": req.Id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("%s %v", nm, err)
	}
	resp := userApi.GetResponse{}
	rows, err := s.Db.Query(query, args...)
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
