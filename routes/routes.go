package routes

import (
	"github.com/MelihEmreGuler/go-user-notes-app/handlers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error { // /
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/signup", handlers.CreateUser) // /signup
	app.Post("/login", handlers.SignIn)      // /login
	app.Post("/logout", handlers.SignOut)    // /logout
	//app.Put("/user/:id", handlers.UpdateUser)    // /user/:id
	//app.Delete("/user/:id", handlers.DeleteUser) // /user/:id
	app.Post("/note", handlers.CreateNote) // /note
	//app.Get("/notes", handlers.GetNotes)         // /notes
	//app.Get("/note/:id", handlers.GetNote)       // /note/:id
	//app.Put("/note/:id", handlers.UpdateNote)    // /note/:id
	//app.Delete("/note/:id", handlers.DeleteNote) // /note/:id
}
