package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v2"
	"time"
)

var Storage *postgres.Storage
var Store *session.Store

func InitConfig() {
	// Initialize custom config
	Storage = postgres.New(postgres.Config{
		DB:            nil,
		ConnectionURI: "postgres://postgres:notes-pass@localhost:5432/note",
		Host:          "localhost",
		Port:          5432,
		Database:      "notes",
		Table:         "fiber_storage",
		SSLMode:       "disable",
		Reset:         false,
		GCInterval:    10 * time.Second,
	})

	// Initialize session store using the config
	Store = session.New(session.Config{
		Storage: Storage,
	})
}

func CreateSession(s *session.Store, c *fiber.Ctx, userId string) error {
	// Get session from storage
	sess, err := s.Get(c)
	if err != nil {
		return errors.New("session not found")
	}

	fmt.Println("set user id to session:", userId)
	sess.Set("user_id", userId)

	fmt.Println("Session id is", sess.ID())

	sess.SetExpiry(5 * time.Minute)
	fmt.Println("session keys:", sess.Keys())

	// Save session
	if err = sess.Save(); err != nil {
		return err
	}

	return nil
}

func GetSession(c *fiber.Ctx) (string, error) {
	sess, err := Store.Get(c)
	if err != nil {
		return "", err
	}

	value := sess.Get("user_id")
	if value == nil {
		return "", errors.New("session key not found")
	}

	return fmt.Sprintf("%v", value), nil
}
