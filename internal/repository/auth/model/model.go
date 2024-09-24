package model

import (
	"database/sql"
	"time"
)

type Role int32

type DeleteRequest struct {
	Id int64 `db:"id"`
}

type GetRequest struct {
	Id int64 `db:"id"`
}

type UpdateRequest struct {
	Id    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type CreateRequest struct {
	Name            string `db:"name"`
	Email           string `db:"email"`
	Password        string `db:"password"`
	PasswordConfirm string `db:"-"`
	Role            Role   `db:"role"`
}

type GetResponse struct {
	Id        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Role      Role         `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
