package repo

import (
	"database/sql"
	"time"
)

type UserAuthToken struct {
	ID        string       `db:"id"`
	Token     string       `db:"token"`
	UserAgent string       `db:"user_agent"`
	CreatedAt time.Time    `db:"create_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
