// Package database utils
package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/feserr/pheme-user/models"
)

// DB database connection
var DB *gorm.DB

// Connect set up the database connection
func Connect(host string, user string, dbname string) {
	dsn := fmt.Sprintf("host=%s user=%s password=docker dbname=%s", host, user, dbname)
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Couldn't connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
