// Package port defines the interfaces (ports) that represent dependencies of the use case layer.
// These ports are implemented by the infrastructure layer and injected into the use cases,
// enabling inversion of control and decoupling business logic from external systems.
package port

import (
	"context"

	"github.com/DucTran999/go-clean-archx/internal/entity"
)

// ProductRepository defines the expected behavior for persisting products.
//
// It is implemented by the infrastructure layer (e.g., database adapter).
// Dependency inversion principle (DIP)
type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
}
