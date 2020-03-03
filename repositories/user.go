package repositories

import (
	"errors"

	"app/models"

	"github.com/jinzhu/gorm"
)

type (
	// UserRepository authentication repository
	UserRepository struct {
		Repository
	}
)

// FindUser from datastore
func (ctl *UserRepository) FindUser(username string, password string) (*models.User, error) {

	model := models.User{}

	if err := ctl.DB(nil).Where("(login is not null and login != '' and login = ?) or (email is not null and login != '' and email = ?)",
		username,
		username,
	).
		First(&model).Error; gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	if !model.ComparePassword(password) {
		// wrong password return like user not found
		return nil, errors.New("Password is invalid!")
	}
	return &model, nil
}
