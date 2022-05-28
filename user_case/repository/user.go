package repository

import (
	"errors"
	"gorm.io/gorm"
	"loginUserGo/entities"
	"loginUserGo/infrastructure/model"
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

func (u UserRepository) Login(data entities.Login) (model.UserModel, error) {
	var entity model.UserModel
	result := u.db.Where("login = ? and password = ?",
		data.Login, data.Password).
		Find(&entity)
	if result.RowsAffected == 0 {
		return entity, errors.New("login or password found")
	}
	return entity, nil
}
