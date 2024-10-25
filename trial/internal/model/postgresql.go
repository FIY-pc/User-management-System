package model

import (
	"User-management-System/trial/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDb *gorm.DB

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
