// Package database utils
package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/feserr/pheme-user/models"
)

// DB database connection
var DB *gorm.DB

// Connect set up the database connection
func Connect() {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPwd := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPwd, dbName)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to the database")
	}

	DB = connection

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic("Couldn't migrate")
	}
}
