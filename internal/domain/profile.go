package domain

type UserProfile struct {
	UserID         uint
	Name           string `json:"name"`
	Email          string `json:"email"`
	LevelCode      string `json:"level_code"`
	LevelName      string `json:"level_name"`
	ProfilePicture string `json:"profile_picture"`
}
