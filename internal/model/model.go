package model

import (
	"database/sql"
	"time"
)

type Role int32

type DeleteRequest struct {
	Id int64
}

type GetRequest struct {
	Id int64
}

type UpdateRequest struct {
	Id    int64
	Name  string
	Email string
}

type CreateRequest struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            Role
}

type GetResponse struct {
	Id        int64
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
