package repo

import (
	"database/sql"
	"time"
)

type UserProfile struct {
	ID                string         `db:"id"`
	UserID            string         `db:"user_id"`
	AboutMe           sql.NullString `db:"about_me"`
	ProfilePhotoCount int            `db:"profile_photo_count"`
	CreatedAt         time.Time      `db:"create_at"`
	UpdatedAt         time.Time      `db:"update_at"`
}
