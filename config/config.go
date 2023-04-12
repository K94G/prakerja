package config

import (
	"prakerja_kg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn := "root:kgxampp@tcp(127.0.0.1:3306)/emergency?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:mfTcjoxxcAZkOtPHIiFy@tcp(containers-us-west-198.railway.app:6659)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:aTeYT6l8CC8PM7oh2S1R@tcp(containers-us-west-10.railway.app:5680)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connected to DB")
	}
	migration()
}

func migration() {
	var user model.User
	var category1 model.Category
	var category2 model.Category

	DB.AutoMigrate(&model.User{})
	DB.First(&user, 1)
	if user.Name == "" {
		user1 := model.User{Id: 1, Name: "Kukuh Gumilang", Username: "kg", Password: "secretkey-kg"}
		DB.Create(&user1)
	}

	DB.AutoMigrate(&model.Category{})
	DB.First(&category1, 1)
	if category1.Name == "" {
		category1 := model.Category{Id: 1, Name: "Pemadam Kebakaran"}
		DB.Create(&category1)
	}
	DB.First(&category2, 2)
	if category2.Name == "" {
		category2 := model.Category{Id: 2, Name: "Ambulance"}
		DB.Create(&category2)
	}

	DB.AutoMigrate(&model.Contact{})
}
