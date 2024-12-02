package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var database *gorm.DB

func DatabaseInit() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	db := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s db=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, db, port)
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "app_data",
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Successfully connect database")
}

func Database() *gorm.DB {
	return database
}
