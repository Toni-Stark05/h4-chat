-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_auth (
                           id TEXT PRIMARY KEY,
                           user_id TEXT NOT NULL REFERENCES users(id),
                           password_hash TEXT NOT NULL,
                           active_auth_count INT DEFAULT 0,
                           create_at TIMESTAMP NOT NULL,
                           deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_auth;
-- +goose StatementEnd
