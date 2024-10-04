-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_profile (
                              id TEXT PRIMARY KEY,
                              user_id TEXT NOT NULL REFERENCES users(id),
                              about_me TEXT,
                              profile_photo_count INT DEFAULT 0,
                              create_at TIMESTAMP NOT NULL,
                              update_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profile;
-- +goose StatementEnd
