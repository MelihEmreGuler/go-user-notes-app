package handlers

import (
	"github.com/MelihEmreGuler/go-user-notes-app/middleware"
	"github.com/MelihEmreGuler/go-user-notes-app/models"
	"github.com/gofiber/fiber/v2"
)

/*curl -X POST -H "Content-Type: application/json" -d '{
  "username": "mertGuler",
  "email": "mert@example.com",
  "password": "mert-password"
}' localhost:8080/signup -v
*/

// CreateUser creates a new user
func CreateUser(c *fiber.Ctx) error {
	type createUserRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	request := &createUserRequest{}
	if err := c.BodyParser(request); err != nil {
		return err
	}
	if err := middleware.SignUp(request.Username, request.Email, request.Password); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true, // Add this line to indicate successful registration
		"message": "user created",
	})
}

/*curl -X POST -H "Content-Type: application/json" -d '{
  "username_or_email": "mertGuler",
  "password": "mert-password"
}' localhost:8080/login -v
*/

func SignIn(c *fiber.Ctx) error {
	var user *models.User
	var err error

	type signInRequest struct {
		UsernameOrEmail string `json:"username_or_email"`
		Password        string `json:"password"`
	}
	request := &signInRequest{}
	if err = c.BodyParser(request); err != nil {
		return err
	}

	// sign in user
	if user, err = middleware.SignIn(request.UsernameOrEmail, request.Password); err != nil {
		return err
	}

	// Create session and store in storage
	if err = middleware.CreateSession(middleware.Store, c, user.ID); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "user logged in",
	})
}

func GetSessionValue(c *fiber.Ctx) error {
	// Get session from storage
	value, err := middleware.GetSession(c)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": value,
	})
}
