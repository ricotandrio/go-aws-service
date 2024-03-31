package configs

import (
	"fmt"
	"os"
	"gorm.io/gorm"
	"go-service-aws/models"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("MYSQL_HOST")
	database, err := gorm.Open(mysql.Open(host), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database")
	} else {
		fmt.Println("Connected to database")
	}

	database.AutoMigrate(&models.Customers{})
	
	DB = database
}

func GetDB() *gorm.DB {
	return DB
}