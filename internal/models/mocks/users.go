package mocks

import (
	"time"

	"example.com/snippetbox/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (models.User, error) {
	if id == 2 {

		user := models.User{}
		user.Name = "dude"
		user.Email = "dupe@example.com"
		user.Created = time.Date(2026, time.February, 26, 0, 0, 0, 0, time.UTC)

		return user, nil
	}
	return models.User{}, models.ErrNoRecord
}

func (m *UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	if id == 1 && currentPassword == "pa$$word" {
		return nil
	}
	return models.ErrInvalidCredentials
}
