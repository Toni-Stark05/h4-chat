package repo

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        string    `db:"id"`
	UserName  string    `db:"user_name"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"create_at"`
	UpdatedAt time.Time `db:"update_at"`
}

func CreateUser(db *sql.DB, user *User) error {
	query := `
    INSERT INTO users (id, user_name, name, email, create_at, update_at)
    VALUES ($1, $2, $3, $4, $5, $6)
    `

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	_, err := db.Exec(query, user.ID, user.UserName, user.Name, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	return nil
}
