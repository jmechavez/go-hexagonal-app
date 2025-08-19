package domain

import (
	"time"
)

// User represents a user entity
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserCreateRequest represents user creation request
type UserCreateRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email"`
}

// UserUpdateRequest represents user update request
type UserUpdateRequest struct {
	ID    int64  `json:"-" validate:"required"`
	Name  string `json:"name" validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email"`
}
