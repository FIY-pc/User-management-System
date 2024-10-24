package model

import (
	"User-management-System/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDb *gorm.DB

func InitPostgres() {
	PostgresDb, err := gorm.Open(postgres.Open(config.Config.Postgres.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := PostgresDb.AutoMigrate(User{}); err != nil {
		panic(err)
	}
	if err := PostgresDb.AutoMigrate(Admin{}); err != nil {
		panic(err)
	}
}
