-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_profile_photos (
                                     user_profile_id TEXT NOT NULL REFERENCES user_profile(id),
                                     profile_photo_id TEXT NOT NULL REFERENCES profile_photo(id),
                                     PRIMARY KEY (user_profile_id, profile_photo_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profile_photos;
-- +goose StatementEnd
