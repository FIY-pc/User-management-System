package model

import (
	"User-management-System/internal/config"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

// Admin 管理员，具有最高权限
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
		hashAdminPass, err := bcrypt.GenerateFromPassword([]byte(rawAdminPass), config.Config.Bcrypt.Cost)
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

// CreateAdmin 创建admin
func CreateAdmin(admin *Admin) error {
	if PostgresDb == nil {
		return errors.New("postgres db is nil")
	}
	resultAdmin := PostgresDb.Model(&Admin{}).Create(admin)
	return resultAdmin.Error
}

// UpdateAdmin 更新admin
func UpdateAdmin(admin *Admin) error {
	if PostgresDb == nil {
		return errors.New("postgres db is nil")
	}
	result := PostgresDb.Model(&Admin{}).Where("admin_name", admin.AdminName).Updates(Admin{AdminName: admin.AdminName, AdminPass: admin.AdminPass})
	return result.Error
}

// GetAdminByName 根据adminName获取admin
func GetAdminByName(AdminName string) (*Admin, error) {
	if PostgresDb == nil {
		return nil, errors.New("postgres db is nil")
	}
	var admin Admin
	resultAdmin := PostgresDb.Where("admin_name = ?", AdminName).First(&admin)
	return &admin, resultAdmin.Error
}

// GetAdminById 根据id获取admin
func GetAdminById(id uint) (*Admin, error) {
	if PostgresDb == nil {
		return nil, errors.New("postgres db is nil")
	}
	var admin Admin
	resultAdmin := PostgresDb.Where("id =?", id).First(&admin)
	return &admin, resultAdmin.Error
}

// DeleteAdminByName 根据adminName删除admin
func DeleteAdminByName(name string) error {
	if PostgresDb == nil {
		return errors.New("postgres db is nil")
	}
	resultAdmin := PostgresDb.Where("admin_name", name).Delete(&Admin{})
	return resultAdmin.Error
}
