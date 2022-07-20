package gorm

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, errors.New("gorm failed to connect")
	}

	db, err := gormDb.DB()
	if err != nil {
		return nil, errors.New("sql failed to connect")
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return gormDb, nil

}
