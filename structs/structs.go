package structs

import (
	"time"
)

type AuthToken struct {
	Token     string
	ExpiresAt time.Time
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type Users struct {
	ID        int64
	Username  string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
