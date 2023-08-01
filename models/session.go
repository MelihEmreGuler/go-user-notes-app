package models

import "time"

type Session struct {
	ID        string
	UserID    string
	LoginTime time.Time
	LastEvent time.Time
	IPAddr    string
	UserAgent string
}
