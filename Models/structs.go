package structs

import (
	"time"
)

type Users struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}
