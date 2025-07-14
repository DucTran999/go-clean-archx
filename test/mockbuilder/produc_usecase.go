// Package mockbuilder provides test builders for mocking dependencies,
// such as repositories and external services.
// It supports fluent-style setup of mock behaviors to simplify unit test configuration.
package mockbuilder

import (
	"context"
	"testing"
	"time"

	"github.com/DucTran999/go-clean-archx/internal/dto"
	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/DucTran999/go-clean-archx/internal/port"
	"github.com/DucTran999/go-clean-archx/test/datatest"
	"github.com/DucTran999/go-clean-archx/test/mocks"
	"github.com/stretchr/testify/mock"
)

// ProductUsecaseBuilder is a builder for setting up mock behaviors for the ProductUsecase interface.
// It provides methods to simulate different scenarios such as successful product creation,
// invalid price errors, and database errors during product creation.
type ProductUsecaseBuilder struct {
	instance *mocks.ProductUsecase
}

// NewProductUsecaseBuilder creates a new ProductUsecaseBuilder instance with a mock ProductUsecase.
// It marks the test as a helper function and returns a builder for configuring product usecase mocks.
func NewProductUsecaseBuilder(t *testing.T) *ProductUsecaseBuilder {
	t.Helper()
	return &ProductUsecaseBuilder{
		instance: mocks.NewProductUsecase(t),
	}
}

// Build returns the mocked ProductUsecase instance for injection into use cases.
func (b *ProductUsecaseBuilder) Build() port.ProductUsecase {
	return b.instance
}

// CreateProductReturnsInvalidPrice sets up the mock to return a domain validation error
// when creating a product. Useful for simulating client-side validation failure.
func (b *ProductUsecaseBuilder) CreateProductReturnsInvalidPrice() *ProductUsecaseBuilder {
	b.instance.EXPECT().
		CreateProduct(mock.Anything, mock.AnythingOfType("dto.CreateProductInput")).
		Return(nil, entity.ErrInvalidPrice)

	return b
}

// CreateProductSuccess sets up the mock to simulate successful product creation, assigning a fixed fake ID.
// Returns the builder for chaining.
func (b *ProductUsecaseBuilder) CreateProductSuccess() *ProductUsecaseBuilder {
	product := entity.Product{
		ID:        datatest.FakeProductID,
		CreatedAt: time.Now(),
	}

	b.instance.EXPECT().
		CreateProduct(mock.Anything, mock.AnythingOfType("dto.CreateProductInput")).
		Run(func(_ context.Context, input dto.CreateProductInput) {
			product.Name = input.Name
			product.Qty = input.Qty
			product.Price = input.Price
		}).
		Return(&product, nil)

	return b
}

// CreateProductErrorDB configures the mock to simulate a database failure during product creation.
func (b *ProductUsecaseBuilder) CreateProductReturnErrDB() *ProductUsecaseBuilder {
	b.instance.EXPECT().
		CreateProduct(mock.Anything, mock.AnythingOfType("dto.CreateProductInput")).
		Return(nil, datatest.ErrUnexpectedDB)

	return b
}
