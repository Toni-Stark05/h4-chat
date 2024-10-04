-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS profile_photo (
                               id TEXT PRIMARY KEY,
                               url TEXT NOT NULL,
                               create_at TIMESTAMP NOT NULL,
                               deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profile_photo;
-- +goose StatementEnd
