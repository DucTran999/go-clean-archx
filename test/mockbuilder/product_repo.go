// Package mockbuilder provides test builders for mocking dependencies,
// such as repositories and external services.
// It supports fluent-style setup of mock behaviors to simplify unit test configuration.
package mockbuilder

import (
	"context"
	"testing"

	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/DucTran999/go-clean-archx/internal/port"
	"github.com/DucTran999/go-clean-archx/test/datatest"
	"github.com/DucTran999/go-clean-archx/test/mocks"
	"github.com/stretchr/testify/mock"
)

// ProductRepoBuilder is a test utility that configures expectations for the ProductRepository mock.
// It follows the builder pattern to streamline setup for different test cases.
type ProductRepoBuilder struct {
	instance *mocks.ProductRepository
}

// NewProductRepoBuilder initializes a new builder with a fresh ProductRepository mock.
// It requires a testing.T object to bind gomock expectations to the test lifecycle.
func NewProductRepoBuilder(t *testing.T) *ProductRepoBuilder {
	t.Helper()
	return &ProductRepoBuilder{
		instance: mocks.NewProductRepository(t),
	}
}

// Build returns the mocked ProductRepository instance for injection into use cases.
func (b *ProductRepoBuilder) Build() port.ProductRepository {
	return b.instance
}

// CreateProductSuccess sets up the mock to simulate successful product creation,
// assigning a fixed fake ID. Returns the builder for chaining.
func (b *ProductRepoBuilder) CreateProductSuccess() *ProductRepoBuilder {
	b.instance.EXPECT().
		Create(mock.Anything, mock.AnythingOfType("*entity.Product")).
		Run(func(_ context.Context, product *entity.Product) {
			product.ID = datatest.FakeProductID
		}).
		Return(nil)

	return b
}

// CreateProductErrorDB configures the mock to simulate a database failure during product creation.
func (b *ProductRepoBuilder) CreateProductErrorDB() *ProductRepoBuilder {
	b.instance.EXPECT().
		Create(mock.Anything, mock.AnythingOfType("*entity.Product")).
		Return(datatest.ErrUnexpectedDB)

	return b
}
