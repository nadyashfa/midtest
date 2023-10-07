package database

import (
	"uts/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "username:password@tcp(localhost:3306)/database_name?parseTime=true"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gagal menghubungkan ke basis data")
	}

	DB.AutoMigrate(&models.User{})
}
