package handlers

import (
	"fmt"
	"github.com/MelihEmreGuler/go-user-notes-app/middleware/authentication"
	"github.com/MelihEmreGuler/go-user-notes-app/middleware/session"
	"github.com/MelihEmreGuler/go-user-notes-app/models"
	"github.com/MelihEmreGuler/go-user-notes-app/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type createUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type signInRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}
type updatePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	SessionID   string `json:"session_id"`
}
type updateEmailRequest struct {
	Password  string `json:"password"`
	Email     string `json:"email"`
	SessionID string `json:"session_id"`
}
type signOutRequest struct {
	SessionID string `json:"session_id"`
}

/*curl -X POST -H "Content-Type: application/json" -d '{
  "username": "mertGuler",
  "email": "mert@example.com",
  "password": "mert-password"
}' localhost:8080/signup -v
*/

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the given username, email, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param request body createUserRequest true "User Information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /signup [post]
func CreateUser(c *fiber.Ctx) error {

	request := &createUserRequest{}
	if err := c.BodyParser(request); err != nil {
		return err
	}
	if err := authentication.SignUp(request.Username, request.Email, request.Password); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true, // Add this line to indicate successful registration
		"message": "User Created",
	})
}

/*curl -X POST -H "Content-Type: application/json" -d '{
  "username_or_email": "mertGuler",
  "password": "mert-password"
}' localhost:8080/login -v
*/

// SignIn signs in a user
// @Summary Sign in a user
// @Description Sign in a user with the given username/email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param request body signInRequest true "User Login Information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /login [post]
func SignIn(c *fiber.Ctx) error {
	var user *models.User
	var err error

	request := &signInRequest{}
	if err = c.BodyParser(request); err != nil {
		return err
	}

	// sign in user
	if user, err = authentication.SignIn(request.UsernameOrEmail, request.Password); err != nil {
		return err
	}

	sessionID := uuid.New().String()
	// Create session and store in database
	if err = session.CreateSession(c, user.ID, sessionID); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":    true, // Add this line to indicate successful login
		"message":    "User successfully logged in",
		"session_id": sessionID,
		"user":       user,
	})
}

// SignOut signs out a user
// @Summary Sign out a user
// @Description Sign out the currently logged-in user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body signOutRequest true "User Logout Information"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /logout [post]
func SignOut(c *fiber.Ctx) error {

	request := &signOutRequest{}
	if err := c.BodyParser(request); err != nil {
		return err
	}

	c.ClearCookie("session_id")
	// Delete session from storage
	if err := session.DeleteSession(c, request.SessionID); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true, // Add this line to indicate successful logout
		"message": "User logged out",
	})
}

// UpdatePassword updates the password of a user
// @Summary Update the password of a user
// @Description Update the password of the currently logged-in user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body updatePasswordRequest true "User Password Information"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user/password [put]
func UpdatePassword(c *fiber.Ctx) error {

	request := &updatePasswordRequest{}

	if err := c.BodyParser(request); err != nil {
		return err
	}

	sessionID := request.SessionID

	sess, err := repository.R.SelectSession(sessionID)
	if err != nil {
		return fmt.Errorf("error while selecting session: %w", err)
	}

	user, err := repository.R.SelectUserById(sess.UserID)
	if err != nil {
		return err
	} else if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	// Update password
	if err = authentication.UpdatePassword(sess.ID, user.Username, request.OldPassword, request.NewPassword); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true, // Add this line to indicate successful password update
		"message": "Password successfully updated",
	})
}

// UpdateEmail updates the email of a user
// @Summary Update the email of a user
// @Description Update the email of the currently logged-in user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body updateEmailRequest true "User Email Information"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user/email [put]
func UpdateEmail(c *fiber.Ctx) error {

	request := &updateEmailRequest{}
	if err := c.BodyParser(request); err != nil {
		return err
	}

	sessionID := request.SessionID

	sess, err := repository.R.SelectSession(sessionID)
	if err != nil {
		return fmt.Errorf("error while selecting session: %w", err)
	}

	// Update email
	if err = authentication.UpdateEmail(sess.UserID, request.Password, request.Email); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true, // Add this line to indicate successful email update
		"message": "Email successfully updated",
	})
}

// DeleteUser deletes a user
// @Summary Delete a user
// @Description Delete the currently logged-in user
// @Tags Users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /user [delete]
func DeleteUser(c *fiber.Ctx) error {
	// Get user from session
	sess, err := session.AuthenticateAndRefresh(c)

	// Delete session from storage
	if err = session.DeleteSession(c, "21323323"); err != nil {
		return err
	} // GECİCİ STRİNG VAR !!
	// Delete user from database
	if err = repository.R.DeleteUser(sess.UserID); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true, // Add this line to indicate successful logout
		"message": "User successfully deleted",
	})
}
