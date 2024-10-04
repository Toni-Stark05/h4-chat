-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_auth_tokens (
                                  id TEXT PRIMARY KEY,
                                  token TEXT NOT NULL,
                                  user_agent TEXT,
                                  create_at TIMESTAMP NOT NULL,
                                  deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_auth_tokens;
-- +goose StatementEnd
