package database

import (
	"fmt"

	"github.com/wasinaphatlilawatthananan/go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// database, err := gorm.Open(mysql.Open("root:keep1234@/go_admin"), &gorm.Config{})
	dsn := "root:keep1234@tcp(localhost:3306)/go_admin?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}
