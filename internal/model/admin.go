package model

import (
	"User-management-System/internal/config"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type Admin struct {
	AdminName string `json:"adminName"`
	AdminPass string `json:"adminPass"`
	gorm.Model
}

// InitAdmin 初始化Admin
// 若admin表为空，根据default配置生成初始admin账号
func InitAdmin() {
	_, err := GetAdminById(1)
	if err != nil {
		initAdmin := Admin{}
		initAdmin.AdminName = config.Config.Admin.AdminName
		rawAdminPass := config.Config.Admin.AdminPass
		hashAdminPass, err := bcrypt.GenerateFromPassword([]byte(rawAdminPass), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		initAdmin.AdminPass = string(hashAdminPass)
		err = CreateAdmin(&initAdmin)
		if err == nil {
			log.Println("InitAdmin fail")
		}
	}
}

func CreateAdmin(admin *Admin) error {
	if PostgresDb == nil {
		return errors.New("postgres db is nil")
	}
	resultAdmin := PostgresDb.Model(&Admin{}).Create(admin)
	return resultAdmin.Error
}

func UpdateAdmin(admin *Admin) error {
	if PostgresDb == nil {
		return errors.New("postgres db is nil")
	}
	resultAdmin := PostgresDb.First(&admin, admin.ID)
	if resultAdmin.Error != nil {
		return resultAdmin.Error
	}
	result := PostgresDb.Save(admin)
	return result.Error
}

func GetAdminByName(AdminName string) (*Admin, error) {
	if PostgresDb == nil {
		return nil, errors.New("postgres db is nil")
	}
	var admin Admin
	resultAdmin := PostgresDb.Where("admin_name = ?", AdminName).First(&admin)
	return &admin, resultAdmin.Error
}

func GetAdminById(id uint) (*Admin, error) {
	if PostgresDb == nil {
		return nil, errors.New("postgres db is nil")
	}
	var admin Admin
	resultAdmin := PostgresDb.Where("id =?", id).First(&admin)
	return &admin, resultAdmin.Error
}

func DeleteAdminByName(name string) error {
	if PostgresDb == nil {
		return errors.New("postgres db is nil")
	}
	resultAdmin := PostgresDb.Where("admin_name=", name).Delete(&Admin{})
	return resultAdmin.Error
}
