// Package usecase contains the business logic implementation of the application.
// It orchestrates operations by coordinating entities and repository interfaces,
// applying business rules and delegating infrastructure work to the appropriate ports.
package usecase

import (
	"context"
	"fmt"

	"github.com/DucTran999/go-clean-archx/internal/dto"
	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/DucTran999/go-clean-archx/internal/port"
)

// productUsecase implements the ProductUsecase interface and handles
// business logic related to product operations.
type productUsecase struct {
	productRepo port.ProductRepository
}

// NewProductUsecase returns a productUsecase instance with the given repository.
func NewProductUsecase(productRepo port.ProductRepository) port.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

// CreateProduct handles the creation of a new product.
func (uc *productUsecase) CreateProduct(ctx context.Context, input dto.CreateProductInput) (*entity.Product, error) {
	product := entity.Product{
		Name:  input.Name,
		Qty:   input.Qty,
		Price: input.Price,
	}

	// Validate domain rules.
	// Although the incoming request is already validated via binding in the controller (e.g., ShouldBindJSON),
	// it's still good practice to validate critical business rules in this layer.
	// This ensures data integrity regardless of the delivery mechanism — such as gRPC, CLI, or tests —
	// which may bypass HTTP-level validation.
	if err := product.IsValid(); err != nil {
		return nil, fmt.Errorf("product is invalid: %w", err)
	}

	// Persist the new product
	if err := uc.productRepo.Create(ctx, &product); err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return &product, nil
}
