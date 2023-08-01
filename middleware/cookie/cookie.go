package cookie

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

const sessionCookieName = "session_id"
const expires = 24 * time.Hour // The session ID can be set to expire after 24 hours

func SetSessionCookie(c *fiber.Ctx, sessionID string) {
	// Create a cookie to send the generated session ID to the user's computer
	cookie := fiber.Cookie{
		Name:     sessionCookieName,
		Value:    sessionID,
		Expires:  time.Now().Add(expires), // Set an expiration time for the cookie
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
}

func GetSessionCookie(c *fiber.Ctx) string {
	// Read the session ID from the cookie stored on the user's computer
	cookie := c.Cookies(sessionCookieName)
	return cookie
}

func DeleteSessionCookie(c *fiber.Ctx) {
	// Invalidate the session ID cookie on the user's computer to delete it
	cookie := fiber.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Set an expired time in the past to invalidate the cookie
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
}
