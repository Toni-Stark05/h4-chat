package repo

import (
	"database/sql"
	"time"
)

type UserAuth struct {
	ID              string       `db:"id"`
	UserID          string       `db:"user_id"`
	PasswordHash    string       `db:"password_hash"`
	ActiveAuthCount int          `db:"active_auth_count"`
	CreatedAt       time.Time    `db:"create_at"`
	DeletedAt       sql.NullTime `db:"deleted_at"`
}
