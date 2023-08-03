package routes

import (
	"github.com/MelihEmreGuler/go-user-notes-app/handlers"
	"github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error { // /
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/swagger/*", swagger.New(swagger.Config{ // /swagger/*
		URL:         "http://localhost:8080/swagger/doc.json", //The url pointing to API definition
		DeepLinking: false,
		Title:       "User Notes API",
	}))

	app.Post("/signup", handlers.CreateUser)           // /signup
	app.Post("/login", handlers.SignIn)                // /login
	app.Post("/logout", handlers.SignOut)              // /logout
	app.Put("/user/password", handlers.UpdatePassword) // /user/password
	app.Put("/user/email", handlers.UpdateEmail)       // /user/email
	app.Delete("/user", handlers.DeleteUser)           // /user
	app.Post("/note", handlers.CreateNote)             // /note
	app.Get("/notes", handlers.GetNotes)               // /notes
	app.Get("/note/:id", handlers.GetNote)             // /note/:id
	app.Put("/note/:id", handlers.UpdateNote)          // /note/:id
	app.Delete("/note/:id", handlers.DeleteNote)       // /note/:id
}
