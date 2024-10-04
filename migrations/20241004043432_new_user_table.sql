-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
                       id TEXT PRIMARY KEY,
                       user_name TEXT NOT NULL,
                       name TEXT,
                       email TEXT NOT NULL,
                       create_at TIMESTAMP NOT NULL,
                       update_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
