package main

import (
	"fmt"
	"log"
	"os"

	"github.com/feserr/pheme-auth/database"
	_ "github.com/feserr/pheme-auth/docs"
	"github.com/feserr/pheme-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
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
	database.Connect()

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

	err := app.Listen(fmt.Sprintf("%v:%v", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))
	if err != nil {
		log.Fatal(err.Error())
	}
}
