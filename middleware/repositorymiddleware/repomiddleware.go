package repositorymiddleware

import (
	"external_api/db"
	"external_api/model"
)

func RepoMiddleware(refreshToken string) (*model.Users, error) {
	db := db.DbManager()
	note := model.Users{}

	err := db.Raw("SELECT id, username, email, password, name, is_login, refresh_token FROM users WHERE refresh_token = ?", refreshToken).Scan(&note).Error
	return &note, err
}
