package db

import (
	"fmt"
	"jdlgj/core"
	"jdlgj/models"

	"github.com/jinzhu/gorm"

	//postgres import
	_ "github.com/jinzhu/gorm/dialects/postgres" //PostgreSQL Driver
)

var db *gorm.DB //database

func init() {

	username := core.GetEnv("PG_USER", "jingyi")
	password := core.GetEnv("PG_PASSWORD", "")
	dbName := core.GetEnv("PG_DATABSE", "jdlgj")
	dbHost := core.GetEnv("PG_HOST", "localhost")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&models.LawFirm{}, &models.Case{}, &models.ServiceCenter{}, &models.Banner{}, &models.Message{}, &models.User{}, &models.WcUser{}) //Database migration
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}
