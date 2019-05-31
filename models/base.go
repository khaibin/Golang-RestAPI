package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func init() {
	godotenv.Load()
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	conn, err := gorm.Open(os.Getenv("db_type"), dbUri)
	if err != nil {
		fmt.Print(err)
	}
	db = conn
	db.Debug().AutoMigrate(&Account{})
}

func GetDB() *gorm.DB {
	return db
}