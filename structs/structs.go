package structs

import (
	"encoding/json"
	"time"
)

type AuthToken struct {
	Token     string
	ExpiresAt time.Time
}

type AuthRequest struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

type Users struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type Post struct {
	UserID        int64  `json:"user_id"`
	Message       string `json:"message"`
	ImageURL      string `json:"image_url"`
	CommunityType string `json:"community_type"`
}

type Profile struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Avatar    string    `json:"avatar"`
	Birthdate time.Time `json:"birthdate"`
	Company   string    `json:"company"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type UserProfile struct {
	Users   `json:"user"`
	Profile `json:"profile"`
}

//Functions

func (up UserProfile) MarshalJSON() ([]byte, error) {
	type Alias UserProfile
	return json.Marshal(&struct {
		UserID    int64     `json:"user_id"`
		Username  string    `json:"username"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Avatar    string    `json:"avatar"`
		Birthdate time.Time `json:"birthdate"`
		Company   string    `json:"company"`
		Gender    string    `json:"gender"`
	}{
		UserID:    int64(up.UserID),
		Username:  up.Username,
		Name:      up.Name,
		Email:     up.Email,
		Avatar:    up.Avatar,
		Birthdate: up.Birthdate,
		Company:   up.Company,
		Gender:    up.Gender,
	})
}
