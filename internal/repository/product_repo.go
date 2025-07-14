// Package repository provides concrete implementations of the repository interfaces.
// It handles data persistence logic, interacting with the database using GORM.
package repository

import (
	"context"

	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/DucTran999/go-clean-archx/internal/port"
	"gorm.io/gorm"
)

// productRepo is the GORM-based implementation of the ProductRepository interface.
type productRepo struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository backed by GORM.
func NewProductRepository(db *gorm.DB) port.ProductRepository {
	return &productRepo{
		db: db,
	}
}

// Create inserts a new product record into the database.
func (r *productRepo) Create(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}
