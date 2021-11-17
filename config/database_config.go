package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnect() *gorm.DB{
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env file")
	}

	dbSource := os.Getenv("DB_SOURCE")

	db, err := gorm.Open(mysql.Open(dbSource), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connect to database")
	}

	return db
}

func CloseDatabaseConnect(db *gorm.DB){
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}