// Package model contains all the models required
// for a functional database management system
package model

import (
	"time"
)

// User model - `users` table
type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	IDAuth    uint64    `json:"-"`
}
