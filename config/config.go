package config

import (
	"prakerja_kg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:kgxampp@tcp(127.0.0.1:3306)/emergency?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:mfTcjoxxcAZkOtPHIiFy@tcp(containers-us-west-198.railway.app:6659)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:aTeYT6l8CC8PM7oh2S1R@tcp(containers-us-west-10.railway.app:5680)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connected to DB")
	}
	migration()
}

func migration() {
	DB.AutoMigrate(&model.Contact{})
}
