package model

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	Name     string
	Password string
	gorm.Model
}

func CreateUser(user *User) error {
	if PostgresDb == nil {
		return errors.New("postgres db not initialized")
	}
	resultUser := PostgresDb.Model(&User{}).Create(user)
	return resultUser.Error
}

func UpdateUser(user *User) error {
	if PostgresDb == nil {
		return errors.New("postgres db not initialized")
	}
	resultUser := PostgresDb.First(&user, user.ID)
	if resultUser.Error != nil {
		return resultUser.Error
	}
	result := PostgresDb.Save(user)
	return result.Error
}

func GetUserByName(Name string) (*User, error) {
	if PostgresDb == nil {
		return nil, errors.New("postgres db not initialized")
	}
	var user User
	resultUser := PostgresDb.Model(&User{}).Where("Name = ?", Name).First(&user)
	return &user, resultUser.Error
}

func DeleteUserByName(Name string) error {
	if PostgresDb == nil {
		return errors.New("postgres db not initialized")
	}
	resultUser := PostgresDb.Delete(&User{Name: Name})
	return resultUser.Error
}
