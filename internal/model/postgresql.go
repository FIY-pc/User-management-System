package model

import (
	"github.com/FIY-pc/User-management-System/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDb 是一个全局的数据库连接对象
var PostgresDb *gorm.DB

// InitPostgres 初始化数据库连接
func InitPostgres() {
	var err error
	PostgresDb, err = gorm.Open(postgres.Open(config.Config.Postgres.Dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	if err = PostgresDb.AutoMigrate(User{}); err != nil {
		panic(err)
	}
	if err = PostgresDb.AutoMigrate(Admin{}); err != nil {
		panic(err)
	}
}
