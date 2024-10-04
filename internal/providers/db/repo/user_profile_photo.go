package repo

type UserProfilePhoto struct {
	UserProfileID  string `db:"user_profile_id"`
	ProfilePhotoID string `db:"profile_photo_id"`
}
