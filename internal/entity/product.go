// Package entity defines domain entities used throughout the business logic layer.
// These entities represent core business objects and should not depend on external libraries
// like frameworks, databases, or transport layers.
package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ErrProductInvalid is a reusable error for any kind of invalid product.
// Instead of defining many specific error variables, we wrap this
// with descriptive context using fmt.Errorf and %w.
// This keeps the error surface small and maintainable,
// supports errors.Is() checks,
// and helps make tests more robust and less brittle.
var ErrProductInvalid = errors.New("invalid product")

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

// IsValid validates the product fields against business rules.
func (p *Product) IsValid() error {
	if p.Name == "" {
		return fmt.Errorf("%w: name cannot be empty", ErrProductInvalid)
	}
	if p.Qty < 0 {
		return fmt.Errorf("%w: quantity must be non-negative", ErrProductInvalid)
	}
	if p.Price <= 0 {
		return fmt.Errorf("%w: price must be greater than zero", ErrProductInvalid)
	}

	return nil
}
