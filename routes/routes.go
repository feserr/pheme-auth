// Package routes file.
package routes

import (
	"github.com/feserr/pheme-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

// Setup creates the routes.
func Setup(app *fiber.App) {
	app.Post("/api/auth/register", controllers.Register)
	app.Post("/api/auth/login", controllers.Login)
	app.Get("/api/auth/user", controllers.User)
	app.Post("/api/auth/logout", controllers.Logout)
	app.Delete("/api/auth/user", controllers.Delete)
}
