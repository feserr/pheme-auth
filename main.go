package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/feserr/pheme-auth/database"
	_ "github.com/feserr/pheme-auth/docs"
	"github.com/feserr/pheme-auth/models"
	"github.com/feserr/pheme-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gopkg.in/yaml.v3"
)

// @title Pheme auth
// @version 1.0
// @description Pheme authentication service.

// @contact.name Elias Serrano
// @contact.email feserr3@gmail.com

// @license.name BSD 2-Clause License
// @license.url http://opensource.org/licenses/BSD-2-Clause

// @BasePath /api/
func main() {
	filename, _ := filepath.Abs("config/network.yml")
	networkFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	networkConfig := models.Network{}
	err = yaml.Unmarshal([]byte(networkFile), &networkConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	dbConfig := networkConfig.Db
	database.Connect(dbConfig.Host, dbConfig.User, dbConfig.Dbname)

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Pheme auth v1.0.0",
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	serverConfig := networkConfig.Server
	app.Listen(fmt.Sprintf("%v:%v", serverConfig.IP, serverConfig.Port))
}
