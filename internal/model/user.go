package model

import (
	"gorm.io/gorm"
)

type User struct {
	Name     string
	Password string
	gorm.Model
}

func createUser(user *User) error {
	resultUser := PostgresDb.Model(&User{}).Create(user)
	return resultUser.Error
}

func UpdateUser(user *User) error {
	resultUser := PostgresDb.First(&user, user.ID)
	if resultUser.Error != nil {
		return resultUser.Error
	}
	result := PostgresDb.Save(user)
	return result.Error
}

func GetUserByName(name string) (*User, error) {
	var user User
	resultUser := PostgresDb.Model(&User{}).Where("name = ?", name).First(&user)
	return &user, resultUser.Error
}

func DeleteUserByName(name string) error {
	resultUser := PostgresDb.Delete(&User{Name: name})
	return resultUser.Error
}
