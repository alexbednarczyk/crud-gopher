package model

import (
	"time"
)

// UserRequest represents metadata of a user
type UserRequest struct {
	DisplayName string `json:"display_name,omitempty" example:"Jonny"`
	FirstName   string `json:"first_name,omitempty" example:"Jon"`
	LastName    string `json:"last_name,omitempty" example:"Doe"`
}

// UserResponse represents metadata of a user
type UserResponse struct {
	ID          int64     `json:"-"`
	GUID        string    `json:"guid"`
	DisplayName string    `json:"display_name,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" `
}
