package storage

import (
	"fmt"
	"time"

	"github.com/ELRAS1/auth/pkg/userApi"
	sq "github.com/Masterminds/squirrel"
)

func (s *Storage) SaveUser(req *userApi.Request) (int64, error) {
	const nm = "[SaveUser]"
	query, sql, err := sq.Insert("users").
		Columns("name", "email", "password", "role", "created_at", "updated_at").
		Values(req.Name, req.Email, req.Password, 0, time.Now(), time.Now()).
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
