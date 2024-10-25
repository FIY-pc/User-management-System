package model

import "gorm.io/gorm"

type Admin struct {
	AdminName string `json:"adminName"`
	AdminPass string `json:"adminPass"`
	gorm.Model
}

func createAdmin(admin *Admin) error {
	resultAdmin := PostgresDb.Model(&Admin{}).Create(admin)
	return resultAdmin.Error
}

func UpdateAdmin(admin *Admin) error {
	resultAdmin := PostgresDb.First(&admin, admin.ID)
	if resultAdmin.Error != nil {
		return resultAdmin.Error
	}
	result := PostgresDb.Save(admin)
	return result.Error
}

func GetAdminByName(AdminName string) (*Admin, error) {
	var admin Admin
	resultAdmin := PostgresDb.Where("AdminName = ?", AdminName).First(&admin)
	return &admin, resultAdmin.Error
}

func DeleteAdminByName(name string) error {
	resultAdmin := PostgresDb.Delete(&Admin{})
	return resultAdmin.Error
}
