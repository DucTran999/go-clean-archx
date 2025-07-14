// Package port defines the interfaces (ports) that represent dependencies of the use case layer.
// These ports are implemented by the infrastructure layer and injected into the use cases,
// enabling inversion of control and decoupling business logic from external systems.
package port

import (
	"context"

	"github.com/DucTran999/go-clean-archx/internal/dto"
	"github.com/DucTran999/go-clean-archx/internal/entity"
)

// ProductUsecase defines the contract for product-related business logic.
//
// The controller depends on this interface to interact with the product repository.
// Dependency inversion principle (DIP)
type ProductUsecase interface {
	CreateProduct(ctx context.Context, input dto.CreateProductInput) (*entity.Product, error)
}
