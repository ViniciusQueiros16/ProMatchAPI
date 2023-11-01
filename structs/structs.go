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
	ID              int64      `json:"id"`
	Username        string     `json:"username"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	UserTypeID      int        `json:"user_type_id"`
	Verified        bool       `json:"verified"`
	PrivacyAccepted bool       `json:"privacy_accepted"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}

type Post struct {
	ID            int    `json:"id"`
	UserID        int64  `json:"user_id"`
	Message       string `json:"message"`
	Image         string `json:"image_url"`
	CommunityType string `json:"community_type"`
	Timestamp     string `json:"timestamp"`
}

type Profile struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Avatar    string    `json:"avatar"`
	Birthdate time.Time `json:"birthdate"`
	Company   string    `json:"company"`
	Gender    string    `json:"gender"`
	About     string    `json:"about"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type Matches struct {
	ID            int  `json:"id"`
	UserID        int  `json:"user_id"`
	MatchedUserID int  `json:"matched_user_id"`
	IsAccepted    bool `json:"is_accepted"`
}

type UserProfile struct {
	Users   `json:"user"`
	Profile `json:"profile"`
}

type ImageRequestBody struct {
	FileName string `json:"filename"`
	Body     string `json:"body"`
}

type ImageUploadResponse struct {
	FileName string `json:"filename"`
	Location string `json:"filelocation"`
}

//Functions

func (up UserProfile) MarshalJSON() ([]byte, error) {
	type Alias UserProfile
	return json.Marshal(&struct {
		UserID   int64  `json:"user_id"`
		Username string `json:"username"`

		Email      string    `json:"email"`
		Avatar     string    `json:"avatar"`
		Birthdate  time.Time `json:"birthdate"`
		Company    string    `json:"company"`
		Gender     string    `json:"gender"`
		UserTypeID int       `json:"user_type_id"`
		Verified   bool      `json:"verified"`
		About      string    `json:"about"`
	}{
		UserID:   int64(up.UserID),
		Username: up.Username,

		Email:      up.Email,
		Avatar:     up.Avatar,
		Birthdate:  up.Birthdate,
		Company:    up.Company,
		Gender:     up.Gender,
		UserTypeID: up.UserTypeID,
		Verified:   up.Verified,
		About:      up.About,
	})
}
