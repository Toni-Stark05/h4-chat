package repo

import (
	"database/sql"
	"time"
)

type ProfilePhoto struct {
	ID        string       `db:"id"`
	URL       string       `db:"url"`
	CreatedAt time.Time    `db:"create_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
