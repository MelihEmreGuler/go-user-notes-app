package main

import (
	"github.com/MelihEmreGuler/go-user-notes-app/database"
	"github.com/MelihEmreGuler/go-user-notes-app/middleware/session"
	"github.com/MelihEmreGuler/go-user-notes-app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

func main() {

	// initialize the database
	database.Init()

	// Postgres Storage for sessions
	session.InitSession()

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// CORS middleware to allow all origins to access our API (for testing purposes)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
	}))

	// Setup routes
	routes.Routes(app)

	// Listen on port 8080
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
