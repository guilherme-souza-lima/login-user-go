package repository

import (
	"gorm.io/gorm"
	"loginUserGo/entities"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u UserRepository) Create(data entities.User) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		err := u.db.Create(&data)
		if err != nil {
			return err.Error
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
