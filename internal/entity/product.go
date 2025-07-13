// Package entity defines domain entities used throughout the business logic layer.
// These entities represent core business objects and should not depend on external libraries
// like frameworks, databases, or transport layers.
package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	// ErrEmptyName is returned when the product name is empty.
	ErrEmptyName = errors.New("product name cannot be empty")

	// ErrQtyNegative is returned when the quantity is less than 0.
	ErrQtyNegative = errors.New("product quantity cannot be negative")

	// ErrInvalidPrice is returned when the product price is not greater than 0.
	ErrInvalidPrice = errors.New("product price must be greater than 0")
)

// Product represents a product in the system with its attributes.
// It is a core domain entity and should be free of infrastructure-specific concerns.
type Product struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	Qty       int        `gorm:"not null;default:0;check:qty >= 0" json:"qty"`
	Price     float64    `gorm:"type:double precision;not null;check:price > 0" json:"price"`
	CreatedAt time.Time  `gorm:"not null;default:now()" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"type:timestamp with time zone;autoUpdateTime" json:"updatedAt,omitempty"`
}

// TableName explicit product schema in database
func (p *Product) TableName() string { return "products" }

// IsValid validates the product fields against business rules.
func (p *Product) IsValid() error {
	if p.Name == "" {
		return ErrEmptyName
	}
	if p.Qty < 0 {
		return ErrQtyNegative
	}
	if p.Price <= 0 {
		return ErrInvalidPrice
	}

	return nil
}
