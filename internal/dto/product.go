// Package dto contains Data Transfer Objects used to pass structured data
// between the delivery layer (e.g., HTTP handlers) and the usecase layer.
package dto

// CreateProductInput represents the input data required to create a new product.
// It is typically populated from a request payload and passed into the usecase.
type CreateProductInput struct {
	Name  string
	Qty   int
	Price float64
}
