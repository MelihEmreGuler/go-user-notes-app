package routes

import (
	"github.com/MelihEmreGuler/go-user-notes-app/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error { // /
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/signup", handlers.CreateUser)      // /signup
	app.Post("/login", handlers.SignIn)           // /login
	app.Get("/session", handlers.GetSessionValue) // /session
}
