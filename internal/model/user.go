package model

import (
	"errors"
	"gorm.io/gorm"
)

// User 用户，具有User权限
type User struct {
	Name     string
	Password string
	gorm.Model
}

// CreateUser 新增用户
func CreateUser(user *User) error {
	if PostgresDb == nil {
		return errors.New("postgres db not initialized")
	}
	resultUser := PostgresDb.Model(&User{}).Create(user)
	return resultUser.Error
}

// UpdateUser 更新用户
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

// GetUserByName 根据用户名获取用户
func GetUserByName(Name string) (*User, error) {
	if PostgresDb == nil {
		return nil, errors.New("postgres db not initialized")
	}
	var user User
	resultUser := PostgresDb.Where("name = ?", Name).First(&user)
	return &user, resultUser.Error
}

// DeleteUserByName 根据用户名删除用户
func DeleteUserByName(Name string) error {
	if PostgresDb == nil {
		return errors.New("postgres db not initialized")
	}
	resultUser := PostgresDb.Delete(&User{Name: Name})
	return resultUser.Error
}
