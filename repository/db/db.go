package db

import (
	"fmt"
	"jdlgj/models"
	"os"

	"github.com/jinzhu/gorm"

	//postgres import
	_ "github.com/jinzhu/gorm/dialects/postgres" //PostgreSQL Driver
)

var db *gorm.DB //database

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func init() {

	username := getEnv("PG_USER", "jingyi")
	password := getEnv("PG_PASSWORD", "")
	dbName := getEnv("PG_DATABSE", "jdlgj")
	dbHost := getEnv("PG_HOST", "localhost")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&models.LawFirm{}, &models.Case{}, &models.ServiceCenter{}, &models.Banner{}, &models.Message{}) //Database migration
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}
