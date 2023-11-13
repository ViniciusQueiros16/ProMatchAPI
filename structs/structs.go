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
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Avatar      string    `json:"avatar"`
	CoverPhoto  string    `json:"cover_photo"`
	PhoneNumber string    `json:"phone_number"`
	Birthdate   time.Time `json:"birthdate"`
	Gender      string    `json:"gender"`
	About       string    `json:"about"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}

type Professional struct {
	ID                     int    `json:"id"`
	UserID                 int    `json:"user_id"`
	TypeService            string `json:"type_service"`
	Recommendations        int    `json:"recommendations"`
	ProfessionalExperience string `json:"professional_experience"`
	LinkSocialMedia        string `json:"link_social_media"`
	LinkPortfolio          string `json:"link_portfolio"`
}

type Client struct {
	ID              int    `json:"id"`
	UserID          int    `json:"user_id"`
	Company         string `json:"company"`
	Recommendations int    `json:"recommendations"`
}

type Matches struct {
	ID            int  `json:"id"`
	UserID        int  `json:"user_id"`
	MatchedUserID int  `json:"matched_user_id"`
	IsAccepted    bool `json:"is_accepted"`
}

type UserProfile struct {
	Users         `json:"user"`
	Profile       `json:"profile"`
	UserAddresses `json:"user_addresses"`
	Notifications `json:"notifications"`
}

type UserAddresses struct {
	Country       string `json:"country"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	State         string `json:"state"`
	PostalCode    string `json:"postal_code"`
}

type Notifications struct {
	Comments          bool   `json:"comments"`
	Candidates        bool   `json:"candidates"`
	Offers            bool   `json:"offers"`
	SMSDeliveryOption string `json:"sms_delivery_option"`
}

type ImageRequestBody struct {
	FileName string `json:"filename"`
	Body     string `json:"body"`
}

type ImageUploadResponse struct {
	FileName string `json:"filename"`
	Location string `json:"filelocation"`
}

type UpdateProfileRequest struct {
	Avatar            string `json:"avatar"`
	Birthdate         string `json:"birthdate"`
	Gender            string `json:"gender"`
	About             string `json:"about"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	CoverPhoto        string `json:"cover_photo"`
	PhoneNumber       string `json:"phone_number"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	UserTypeID        int    `json:"user_type_id"`
	PrivacyAccepted   bool   `json:"privacy_accepted"`
	Country           string `json:"country"`
	StreetAddress     string `json:"street_address"`
	City              string `json:"city"`
	State             string `json:"state"`
	PostalCode        string `json:"postal_code"`
	Comments          bool   `json:"comments"`
	Candidates        bool   `json:"candidates"`
	Offers            bool   `json:"offers"`
	SMSDeliveryOption string `json:"sms_delivery_option"`
	Company           string `json:"company"`
	TypeService       string `json:"type_service"`
}

//Functions

func (up UserProfile) MarshalJSON() ([]byte, error) {
	type Alias UserProfile
	return json.Marshal(&struct {
		UserID          int64         `json:"user_id"`
		Username        string        `json:"username"`
		Email           string        `json:"email"`
		Avatar          string        `json:"avatar"`
		Birthdate       time.Time     `json:"birthdate"`
		Gender          string        `json:"gender"`
		UserTypeID      int           `json:"user_type_id"`
		Verified        bool          `json:"verified"`
		PrivacyAccepted bool          `json:"privacy_accepted"`
		About           string        `json:"about"`
		FirstName       string        `json:"first_name"`
		LastName        string        `json:"last_name"`
		CoverPhoto      string        `json:"cover_photo"`
		PhoneNumber     string        `json:"phone_number"`
		UserAddresses   UserAddresses `json:"user_addresses"`
		Notifications   Notifications `json:"notifications"`
	}{
		UserID:          int64(up.UserID),
		Username:        up.Username,
		Email:           up.Email,
		Avatar:          up.Avatar,
		Birthdate:       up.Birthdate,
		Gender:          up.Gender,
		UserTypeID:      up.UserTypeID,
		Verified:        up.Verified,
		PrivacyAccepted: up.PrivacyAccepted,
		About:           up.About,
		FirstName:       up.FirstName,
		LastName:        up.LastName,
		CoverPhoto:      up.CoverPhoto,
		PhoneNumber:     up.PhoneNumber,
		UserAddresses:   up.UserAddresses,
		Notifications:   up.Notifications,
	})
}
