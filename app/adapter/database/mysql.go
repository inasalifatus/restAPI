package database

import (
	"restAPI/app/models"
	"restAPI/app/shared/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	DB *gorm.DB
}

func SetupDatabase() *MySQL {

	url := config.DB.MySQL.DataSourceName
	db, err := gorm.Open(mysql.Open(url))

	db.AutoMigrate(&models.Order{}, &models.Item{})

	if err != nil {
		panic("")
	}

	Db, err := db.DB()

	if err != nil {
		panic("")
	}

	Db.Ping()

	return &MySQL{
		DB: db,
	}
}
