// Package routes file.
package routes

import (
	"github.com/feserr/pheme-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setup creates the routes.
func Setup(app *fiber.App) {
	app.Post("/api/v1/auth/register", controllers.Register)
	app.Post("/api/v1/auth/login", controllers.Login)
	app.Get("/api/v1/auth/user", controllers.User)
	app.Post("/api/v1/auth/logout", controllers.Logout)
	app.Delete("/api/v1/auth/user", controllers.Delete)
}
