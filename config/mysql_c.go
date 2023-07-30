package config

import (
	"blog-app/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() {
	dsn := "root:password@tcp(127.0.0.1:3306)/blog-app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[DATABASE] MYSQL Connected")
	db.AutoMigrate(&model.Blog{})
}
